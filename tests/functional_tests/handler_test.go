// +build all handler

package eventing

import (
	"fmt"
	"testing"
	"time"
)

func TestN1QLLabelledBreak(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_labelled_break.js"

	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)

	setIndexStorageMode()
	time.Sleep(time.Second * 5)
	fireQuery("CREATE PRIMARY INDEX on default;")
	time.Sleep(time.Second * 5)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "N1QLLabelledBreak",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)
}

func TestN1QLUnlabelledBreak(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_unlabelled_break.js"

	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)

	setIndexStorageMode()
	time.Sleep(time.Second * 5)
	fireQuery("CREATE PRIMARY INDEX on default;")
	time.Sleep(time.Second * 5)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "N1QLUnlabelledBreak",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)
}

func TestN1QLThrowStatement(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_throw_statement.js"

	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)

	setIndexStorageMode()
	time.Sleep(time.Second * 5)
	fireQuery("CREATE PRIMARY INDEX on default;")
	time.Sleep(time.Second * 5)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "N1QLThrowStatement",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)
}

func TestN1QLNestedForLoop(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_nested_for_loops.js"

	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)

	setIndexStorageMode()
	time.Sleep(time.Second * 5)
	fireQuery("CREATE PRIMARY INDEX on default;")
	time.Sleep(time.Second * 5)
	createAndDeployFunction(handler, handler, &commonSettings{lcbInstCap: 6})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "N1QLNestedForLoop",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	fireQuery("DROP PRIMARY INDEX on default;")
	flushFunctionAndBucket(handler)
}

func TestOnUpdateBucketOpDefaultSettings(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_update.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "OnUpdateBucketOpDefaultSettings",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestOnUpdateBucketOpNonDefaultSettings(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_update.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{thrCount: 4, batchSize: 77})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "OnUpdateBucketOpNonDefaultSettings",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestOnUpdateN1QLOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_insert_on_update.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "OnUpdateN1QLOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestOnDeleteBucketOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_delete.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{expiry: 1, delete: true}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "OnDeleteBucketOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestDocTimerBucketOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_with_doc_timer.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "DocTimerBucketOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestDocTimerN1QLOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_insert_with_doc_timer.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "DocTimerN1QLOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestCronTimerBucketOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_with_cron_timer.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "CronTimerBucketOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestCronTimerN1QLOp(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "n1ql_insert_with_cron_timer.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "CronTimerN1QLOp",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}

func TestDeployUndeployLoopDefaultSettings(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_update.js"
	flushFunctionAndBucket(handler)

	for i := 0; i < 5; i++ {
		createAndDeployFunction(handler, handler, &commonSettings{})

		pumpBucketOps(opsType{}, &rateLimit{})
		eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
		if itemCount != eventCount {
			t.Error("For", "DeployUndeployLoopDefaultSettings",
				"expected", itemCount,
				"got", eventCount,
			)
		}

		dumpStats(handler)
		fmt.Println("Undeploying app:", handler)
		setSettings(handler, false, false, &commonSettings{})
		bucketFlush("default")
		bucketFlush("hello-world")
		time.Sleep(30 * time.Second)
	}

	deleteFunction(handler)
}

func TestDeployUndeployLoopDocTimer(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_with_doc_timer.js"
	flushFunctionAndBucket(handler)

	for i := 0; i < 5; i++ {
		createAndDeployFunction(handler, handler, &commonSettings{})

		pumpBucketOps(opsType{}, &rateLimit{})
		eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
		if itemCount != eventCount {
			t.Error("For", "DeployUndeployLoopDocTimer",
				"expected", itemCount,
				"got", eventCount,
			)
		}

		dumpStats(handler)
		fmt.Println("Undeploying app:", handler)
		setSettings(handler, false, false, &commonSettings{})
		bucketFlush("default")
		bucketFlush("hello-world")
		time.Sleep(30 * time.Second)
	}

	deleteFunction(handler)
}

func TestDeployUndeployLoopNonDefaultSettings(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_update.js"
	flushFunctionAndBucket(handler)

	for i := 0; i < 5; i++ {
		createAndDeployFunction(handler, handler, &commonSettings{thrCount: 4, batchSize: 77})

		pumpBucketOps(opsType{}, &rateLimit{})
		eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
		if itemCount != eventCount {
			t.Error("For", "DeployUndeployLoopNonDefaultSettings",
				"expected", itemCount,
				"got", eventCount,
			)
		}

		dumpStats(handler)
		fmt.Println("Undeploying app:", handler)
		setSettings(handler, false, false, &commonSettings{})
		bucketFlush("default")
		bucketFlush("hello-world")
		time.Sleep(30 * time.Second)
	}

	deleteFunction(handler)
}

func TestMultipleHandlers(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler1 := "bucket_op_on_update.js"
	handler2 := "n1ql_insert_on_update.js"

	flushFunctionAndBucket(handler1)
	flushFunctionAndBucket(handler2)

	createAndDeployFunction(handler1, handler1, &commonSettings{})
	createAndDeployFunction(handler2, handler2, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount*2, statsLookupRetryCounter*2)
	if itemCount*2 != eventCount {
		t.Error("For", "MultipleHandlers",
			"expected", itemCount*2,
			"got", eventCount,
		)
	}

	dumpStats(handler1)
	dumpStats(handler2)

	// Pause the apps
	setSettings(handler1, true, false, &commonSettings{})
	setSettings(handler2, true, false, &commonSettings{})

	flushFunctionAndBucket(handler1)
	flushFunctionAndBucket(handler2)
}

func TestPauseResumeLoopDefaultSettings(t *testing.T) {
	time.Sleep(5 * time.Second)

	handler := "bucket_op_on_update.js"

	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{})

	for i := 0; i < 5; i++ {
		if i > 0 {
			setSettings(handler, true, true, &commonSettings{})
		}

		pumpBucketOps(opsType{startIndex: itemCount * i}, &rateLimit{})
		eventCount := verifyBucketOps(itemCount*(i+1), statsLookupRetryCounter)
		if itemCount*(i+1) != eventCount {
			t.Error("For", "PauseAndResumeLoopDefaultSettings",
				"expected", itemCount*(i+1),
				"got", eventCount,
			)
		}

		dumpStats(handler)
		fmt.Printf("Pausing the app: %s\n\n", handler)
		setSettings(handler, true, false, &commonSettings{})
	}

	flushFunctionAndBucket(handler)
}

func TestPauseResumeLoopNonDefaultSettings(t *testing.T) {
	time.Sleep(5 * time.Second)

	handler := "bucket_op_on_update.js"

	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{thrCount: 4, batchSize: 77, workerCount: 4})

	for i := 0; i < 5; i++ {
		if i > 0 {
			setSettings(handler, true, true, &commonSettings{thrCount: 4, batchSize: 77, workerCount: 4})
		}

		pumpBucketOps(opsType{startIndex: itemCount * i}, &rateLimit{})
		eventCount := verifyBucketOps(itemCount*(i+1), statsLookupRetryCounter)
		if itemCount*(i+1) != eventCount {
			t.Error("For", "PauseAndResumeLoopNonDefaultSettings",
				"expected", itemCount*(i+1),
				"got", eventCount,
			)
		}

		dumpStats(handler)
		fmt.Printf("Pausing the app: %s\n\n", handler)
		setSettings(handler, true, false, &commonSettings{})
	}

	flushFunctionAndBucket(handler)
}

func TestCommentUnCommentOnDelete(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "on_delete_bucket_op_comment.js"
	appName := "comment_uncomment_test"
	flushFunctionAndBucket(handler)

	createAndDeployFunction(appName, handler, &commonSettings{})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "CommentUnCommentOnDelete",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(appName)
	fmt.Println("Undeploying app:", appName)
	setSettings(appName, false, false, &commonSettings{})

	time.Sleep(30 * time.Second)

	handler = "on_delete_bucket_op_uncomment.js"
	createAndDeployFunction(appName, handler, &commonSettings{})

	pumpBucketOps(opsType{delete: true}, &rateLimit{})
	eventCount = verifyBucketOps(0, statsLookupRetryCounter)
	if eventCount != 0 {
		t.Error("For", "CommentUnCommentOnDelete",
			"expected", 0,
			"got", eventCount,
		)
	}

	dumpStats(appName)
	fmt.Println("Undeploying app:", appName)
	setSettings(appName, false, false, &commonSettings{})

	time.Sleep(5 * time.Second)
	flushFunctionAndBucket(handler)
}

func TestCPPWorkerCleanup(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "bucket_op_on_update.js"
	flushFunctionAndBucket(handler)
	createAndDeployFunction(handler, handler, &commonSettings{batchSize: 100, workerCount: 16})

	pumpBucketOps(opsType{}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "CPPWorkerCleanup",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
	time.Sleep(30 * time.Second)
}

func TestWithUserXattrs(t *testing.T) {
	time.Sleep(time.Second * 5)
	handler := "on_delete_bucket_op_comment.js"
	flushFunctionAndBucket(handler)
	pumpBucketOps(opsType{}, &rateLimit{})
	createAndDeployFunction(handler, handler, &commonSettings{streamBoundary: "from_now"})
	waitForDeployToFinish(handler)

	pumpBucketOps(opsType{writeXattrs: true}, &rateLimit{})
	eventCount := verifyBucketOps(itemCount, statsLookupRetryCounter)
	if itemCount != eventCount {
		t.Error("For", "WithUserXattrs",
			"expected", itemCount,
			"got", eventCount,
		)
	}

	dumpStats(handler)
	flushFunctionAndBucket(handler)
}
