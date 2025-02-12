package test

import (
	"testing"

	"github.com/openfga/openfga/storage"
)

func RunAllTests(t *testing.T, ds storage.OpenFGADatastore) {
	RunQueryTests(t, ds)
	RunCommandTests(t, ds)
}

func RunQueryTests(t *testing.T, ds storage.OpenFGADatastore) {
	t.Run("TestCheckQuery", func(t *testing.T) { TestCheckQuery(t, ds) })
	t.Run("TestCheckQueryAgainstGitHubModel", func(t *testing.T) { TestCheckQueryAgainstGitHubModel(t, ds) })
	t.Run("TestCheckQueryWithContextualTuplesAgainstGitHubModel", func(t *testing.T) { TestCheckQueryWithContextualTuplesAgainstGitHubModel(t, ds) })
	t.Run("TestCheckQueryAuthorizationModelsVersioning", func(t *testing.T) { TestCheckQueryAuthorizationModelsVersioning(t, ds) })

	t.Run("TestReadAuthorizationModelQueryErrors", func(t *testing.T) { TestReadAuthorizationModelQueryErrors(t, ds) })
	t.Run("TestSuccessfulReadAuthorizationModelQuery", func(t *testing.T) { TestSuccessfulReadAuthorizationModelQuery(t, ds) })
	t.Run("TestReadAuthorizationModel", func(t *testing.T) { ReadAuthorizationModelTest(t, ds) })

	t.Run("TestExpandQuery", func(t *testing.T) { TestExpandQuery(t, ds) })
	t.Run("TestExpandQueryErrors", func(t *testing.T) { TestExpandQueryErrors(t, ds) })

	t.Run("TestGetStoreQuery", func(t *testing.T) { TestGetStoreQuery(t, ds) })
	t.Run("TestGetStoreSucceeds", func(t *testing.T) { TestGetStoreSucceeds(t, ds) })
	t.Run("TestListStores", func(t *testing.T) { TestListStores(t, ds) })

	t.Run("TestReadAssertionQuery", func(t *testing.T) { TestReadAssertionQuery(t, ds) })

	t.Run("TestReadQuery", func(t *testing.T) { TestReadQuery(t, ds) })

	t.Run("TestReadTuplesQuery", func(t *testing.T) { TestReadTuplesQuery(t, ds) })
	t.Run("TestReadTuplesQueryInvalidContinuationToken",
		func(t *testing.T) { TestReadTuplesQueryInvalidContinuationToken(t, ds) },
	)

	t.Run("TestReadAuthorizationModelsWithoutPaging",
		func(t *testing.T) { TestReadAuthorizationModelsWithoutPaging(t, ds) },
	)

	t.Run("TestReadAuthorizationModelsWithPaging",
		func(t *testing.T) { TestReadAuthorizationModelsWithPaging(t, ds) },
	)

	t.Run("TestReadAuthorizationModelsInvalidContinuationToken",
		func(t *testing.T) { TestReadAuthorizationModelsInvalidContinuationToken(t, ds) },
	)

	t.Run("TestReadChanges", func(t *testing.T) { TestReadChanges(t, ds) })
	t.Run("TestReadChangesReturnsSameContTokenWhenNoChanges",
		func(t *testing.T) { TestReadChangesReturnsSameContTokenWhenNoChanges(t, ds) },
	)
	t.Run("TestListObjects", func(t *testing.T) { TestListObjects(t, ds) })
}

func RunCommandTests(t *testing.T, ds storage.OpenFGADatastore) {
	t.Run("TestWriteCommand", func(t *testing.T) { TestWriteCommand(t, ds) })
	t.Run("TestWriteAuthorizationModel", func(t *testing.T) { WriteAuthorizationModelTest(t, ds) })
	t.Run("TestWriteAssertions", func(t *testing.T) { TestWriteAssertions(t, ds) })
	t.Run("TestCreateStore", func(t *testing.T) { TestCreateStore(t, ds) })
	t.Run("TestDeleteStore", func(t *testing.T) { TestDeleteStore(t, ds) })
}

func RunAllBenchmarks(b *testing.B, ds storage.OpenFGADatastore) {
	RunCheckBenchmarks(b, ds)
}

func RunCheckBenchmarks(b *testing.B, ds storage.OpenFGADatastore) {
	b.Run("BenchmarkCheckWithoutTrace", func(b *testing.B) { BenchmarkCheckWithoutTrace(b, ds) })
	b.Run("BenchmarkWithTrace", func(b *testing.B) { BenchmarkWithTrace(b, ds) })
}
