/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package template

// annotation keys
const (
	// IconClassAnnotation is the rss class of an icon representing a template
	IconClassAnnotation = "iconClass"

	// ServiceBrokerRoot is the API root of the template service broker.
	ServiceBrokerRoot = "/brokers/template.openshift.io"

	// ServiceMetadataIconClass is the key for the template iconClass as returned
	// in the services.metadata map from a service broker catalog response
	ServiceMetadataIconClass = "console.openshift.io/iconClass"

	// TemplateUIDIndex is the name of an index on the generated template lister,
	// initialised and used by the template service broker.
	TemplateUIDIndex = "templateuid"

	// ExposeAnnotationPrefix indicates that part of an object in a template
	// should be exposed in some way, for example implying that it should be
	// returned by the template service broker in the results of a bind call.
	// The rest of the annotation name following the prefix may be used by the
	// exposer as a key name.  The annotation value is a Kubernetes JSONPath
	// template expression which the exposer uses to calculate the exposed
	// value.  JSONPath expressions which return multiple and/or complex objects
	// are not permitted (with the exception of []byte, which is permitted).
	// Any []byte values returned are converted to strings.
	ExposeAnnotationPrefix = "template.openshift.io/expose-"

	// Base64ExposeAnnotationPrefix is as ExposeAnnotationPrefix, except that
	// any []byte values returned are base64 encoded.
	Base64ExposeAnnotationPrefix = "template.openshift.io/base64-expose-"

	// WaitForReadyAnnotation indicates that the TemplateInstance controller
	// should wait for the object to be ready before reporting the template
	// instantiation complete.
	WaitForReadyAnnotation = "template.alpha.openshift.io/wait-for-ready"

	// BindableAnnotation indicates whether the template service broker should
	// advertise the template as being bindable (default is true)
	BindableAnnotation = "template.openshift.io/bindable"
)
