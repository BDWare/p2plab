// Copyright 2019 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metadata

import bolt "go.etcd.io/bbolt"

var (
	// API Resources.
	bucketKeyVersion     = []byte(schemaVersion)
	bucketKeyClusters    = []byte("clusters")
	bucketKeyNodes       = []byte("nodes")
	bucketKeyScenarios   = []byte("scenarios")
	bucketKeyBuilds      = []byte("builds")
	bucketKeyBenchmarks  = []byte("benchmarks")
	bucketKeyExperiments = []byte("experiments")

	// Cluster buckets.
	bucketKeySize         = []byte("size")
	bucketKeyInstanceType = []byte("instanceType")
	bucketKeyRegion       = []byte("region")

	// Scenario buckets.
	bucketKeyObjects   = []byte("objects")
	bucketKeySeed      = []byte("seed")
	bucketKeyBenchmark = []byte("benchmark")
	bucketKeyType      = []byte("type")
	bucketKeySource    = []byte("source")
	bucketKeyLayout    = []byte("layout")
	bucketKeyChunker   = []byte("chunker")
	bucketKeyRawLeaves = []byte("rawLeaves")
	bucketKeyHashFunc  = []byte("hashFunc")
	bucketKeyMaxLinks  = []byte("maxLinks")

	// Node buckets.
	bucketKeyAddress            = []byte("address")
	bucketKeyAgentPort          = []byte("agentPort")
	bucketKeyAppPort            = []byte("appPort")
	bucketKeyPort               = []byte("port")
	bucketKeyTransports         = []byte("transports")
	bucketKeyMuxers             = []byte("muxers")
	bucketKeySecurityTransports = []byte("securityTransports")
	bucketKeyRouting            = []byte("routing")

	// Build buckets
	bucketKeyLink = []byte("link")

	// Benchmark buckets.
	bucketKeyCluster  = []byte("cluster")
	bucketKeyScenario = []byte("scenario")
	bucketKeyPlan     = []byte("plan")
	bucketKeySubject  = []byte("subject")
	bucketKeyReport   = []byte("report")

	// Common buckets.
	bucketKeyID           = []byte("id")
	bucketKeyStatus       = []byte("status")
	bucketKeyLabels       = []byte("labels")
	bucketKeyCreatedAt    = []byte("createdAt")
	bucketKeyUpdatedAt    = []byte("updatedAt")
	bucketKeyDefinition   = []byte("definition")
	bucketKeyGitReference = []byte("gitReference")
)

func getBucket(tx *bolt.Tx, keys ...[]byte) *bolt.Bucket {
	bkt := tx.Bucket(keys[0])

	for _, key := range keys[1:] {
		if bkt == nil {
			break
		}
		bkt = bkt.Bucket(key)
	}

	return bkt
}

func createBucketIfNotExists(tx *bolt.Tx, keys ...[]byte) (*bolt.Bucket, error) {
	bkt, err := tx.CreateBucketIfNotExists(keys[0])
	if err != nil {
		return nil, err
	}

	for _, key := range keys[1:] {
		bkt, err = bkt.CreateBucketIfNotExists(key)
		if err != nil {
			return nil, err
		}
	}

	return bkt, nil
}

func getClustersBucket(tx *bolt.Tx) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyClusters)
}

func getClusterBucket(tx *bolt.Tx, id string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyClusters, []byte(id))
}

func createClustersBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyClusters)
}

func getNodesBucket(tx *bolt.Tx, cluster string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyClusters, []byte(cluster), bucketKeyNodes)
}

func getNodeBucket(tx *bolt.Tx, cluster, id string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyClusters, []byte(cluster), bucketKeyNodes, []byte(id))
}

func createNodesBucket(tx *bolt.Tx, cluster string) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyClusters, []byte(cluster), bucketKeyNodes)
}

func getScenariosBucket(tx *bolt.Tx) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyScenarios)
}

func getScenarioBucket(tx *bolt.Tx, name string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyScenarios, []byte(name))
}

func createScenariosBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyScenarios)
}

func getBuildsBucket(tx *bolt.Tx) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyBuilds)
}

func getBuildBucket(tx *bolt.Tx, id string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyBuilds, []byte(id))
}

func createBuildsBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyBuilds)
}

func getBenchmarksBucket(tx *bolt.Tx) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyBenchmarks)
}

func getBenchmarkBucket(tx *bolt.Tx, id string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyBenchmarks, []byte(id))
}

func createBenchmarksBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyBenchmarks)
}

func getExperimentsBucket(tx *bolt.Tx) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyExperiments)
}

func getExperimentBucket(tx *bolt.Tx, name string) *bolt.Bucket {
	return getBucket(tx, bucketKeyVersion, bucketKeyExperiments, []byte(name))
}

func createExperimentsBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return createBucketIfNotExists(tx, bucketKeyVersion, bucketKeyExperiments)
}
