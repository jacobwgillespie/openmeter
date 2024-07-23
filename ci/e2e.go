package main

import (
	"context"
	"fmt"

	"github.com/openmeterio/openmeter/ci/internal/dagger"
)

func (m *Ci) Etoe(
	// +optional
	test string,
) *Etoe {
	return &Etoe{
		Ci: m,
	}
}

type Etoe struct {
	Ci *Ci
}

// func (e *Etoe) All(ctx context.Context) error {
// 	p := pool.New().WithErrors().WithContext(ctx)

// 	p.Go(syncFunc(e.App("")))
// 	p.Go(syncFunc(e.Diff("HEAD")))

// 	return p.Wait()
// }

func (e *Etoe) App(
	// +optional
	test string,
) *dagger.Container {
	image := e.Ci.Build().ContainerImage("").
		WithExposedPort(10000).
		WithMountedFile("/etc/openmeter/config.yaml", e.Ci.Source.File("e2e/config.yaml")).
		WithServiceBinding("kafka", dag.Kafka(dagger.KafkaOpts{Version: kafkaVersion}).Service()).
		WithServiceBinding("clickhouse", clickhouse())

	api := image.
		WithExposedPort(8080).
		WithServiceBinding("postgres", postgres()).
		WithEnvVariable("POSTGRES_HOST", "postgres").
		WithExec([]string{"openmeter", "--config", "/etc/openmeter/config.yaml"}).
		AsService()

	sinkWorker := image.
		WithServiceBinding("redis", redis()).
		WithServiceBinding("api", api). // Make sure api is up before starting sink worker
		WithExec([]string{"openmeter-sink-worker", "--config", "/etc/openmeter/config.yaml"}).
		AsService()

	args := []string{"go", "test", "-v"}

	if test != "" {
		args = append(args, "-run", fmt.Sprintf("Test%s", test))
	}

	args = append(args, "./e2e/app/...")

	return dag.Go(dagger.GoOpts{
		Container: goModule().
			WithSource(e.Ci.Source).
			Container().
			WithServiceBinding("api", api).
			WithServiceBinding("sink-worker", sinkWorker).
			WithEnvVariable("OPENMETER_ADDRESS", "http://api:8080").
			WithEnvVariable("TEST_WAIT_ON_START", "true"),
	}).
		Exec(args)
}

func (e *Etoe) Diff(
	ctx context.Context,
	// The base ref against which to run the diffs
	baseRef string,
) error {
	db := postgres()

	// Migrate DB with base state
	_, err := goModule().
		WithCgoEnabled().
		WithSource(e.Ci.Source).
		Container().
		WithExec([]string{"apk", "add", "--update", "--no-cache", "ca-certificates", "make", "git", "curl", "clang", "lld"}).
		WithExec([]string{"git", "checkout", baseRef}).
		WithServiceBinding("postgres", db).
		WithExec([]string{"go", "run", "./tools/migrate"}).
		Sync(ctx)
	if err != nil {
		return fmt.Errorf("failed to migrate base ref %s: %w", baseRef, err)
	}

	_, err = goModule().
		WithSource(e.Ci.Source).
		Container().
		WithServiceBinding("postgres", db).
		WithExec([]string{"go", "test", "-v", "./e2e/diff/..."}).
		Sync(ctx)

	return err
}

func clickhouse() *dagger.Service {
	return dag.Container().
		From(fmt.Sprintf("clickhouse/clickhouse-server:%s-alpine", clickhouseVersion)).
		WithEnvVariable("CLICKHOUSE_USER", "default").
		WithEnvVariable("CLICKHOUSE_PASSWORD", "default").
		WithEnvVariable("CLICKHOUSE_DB", "openmeter").
		WithExposedPort(9000).
		WithExposedPort(9009).
		WithExposedPort(8123).
		AsService()
}

func redis() *dagger.Service {
	return dag.Container().
		From(fmt.Sprintf("redis:%s-alpine", redisVersion)).
		WithExposedPort(6379).
		AsService()
}

func postgres() *dagger.Service {
	return dag.Container().
		From(fmt.Sprintf("postgres:%s", postgresVersion)).
		WithEnvVariable("POSTGRES_USER", "postgres").
		WithEnvVariable("POSTGRES_PASSWORD", "postgres").
		WithEnvVariable("POSTGRES_DB", "postgres").
		WithExposedPort(5432).
		AsService()
}
