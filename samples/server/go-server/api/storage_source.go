/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// StorageSource describes the location of the source in an archive file in Google Cloud Storage.
type StorageSource struct {

	// Google Cloud Storage bucket containing source (see [Bucket Name Requirements] (https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	Bucket string `json:"bucket,omitempty"`

	// Google Cloud Storage object containing source.
	Object string `json:"object,omitempty"`

	// Google Cloud Storage generation for the object.
	Generation string `json:"generation,omitempty"`
}