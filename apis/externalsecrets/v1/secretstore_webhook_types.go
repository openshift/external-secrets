/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"
)

// WebHookProvider Configures an store to sync secrets from simple web apis.
type WebhookProvider struct {
	// Webhook Method
	// +optional, default GET
	Method string `json:"method,omitempty"`

	// Webhook url to call
	URL string `json:"url"`

	// Headers
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	// Auth specifies a authorization protocol. Only one protocol may be set.
	// +optional
	Auth *AuthorizationProtocol `json:"auth,omitempty"`

	// Body
	// +optional
	Body string `json:"body,omitempty"`

	// Timeout
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Result formatting
	Result WebhookResult `json:"result"`

	// Secrets to fill in templates
	// These secrets will be passed to the templating function as key value pairs under the given name
	// +optional
	Secrets []WebhookSecret `json:"secrets,omitempty"`

	// PEM encoded CA bundle used to validate webhook server certificate. Only used
	// if the Server URL is using HTTPS protocol. This parameter is ignored for
	// plain HTTP protocol connection. If not set the system root certificates
	// are used to validate the TLS connection.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`

	// The provider for the CA bundle to use to validate webhook server certificate.
	// +optional
	CAProvider *WebhookCAProvider `json:"caProvider,omitempty"`
}

// AuthorizationProtocol contains the protocol-specific configuration
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:MaxProperties=1
type AuthorizationProtocol struct {
	// NTLMProtocol configures the store to use NTLM for auth
	// +optional
	NTLM *NTLMProtocol `json:"ntlm"`

	// Define other protocols here
}

// NTLMProtocol contains the NTLM-specific configuration.
type NTLMProtocol struct {
	UserName esmeta.SecretKeySelector `json:"usernameSecret"`
	Password esmeta.SecretKeySelector `json:"passwordSecret"`
}
type WebhookCAProviderType string

const (
	WebhookCAProviderTypeSecret    WebhookCAProviderType = "Secret"
	WebhookCAProviderTypeConfigMap WebhookCAProviderType = "ConfigMap"
)

// Defines a location to fetch the cert for the webhook provider from.
type WebhookCAProvider struct {
	// The type of provider to use such as "Secret", or "ConfigMap".
	// +kubebuilder:validation:Enum="Secret";"ConfigMap"
	Type WebhookCAProviderType `json:"type"`

	// The name of the object located at the provider type.
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MaxLength:=253
	// +kubebuilder:validation:Pattern:=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$
	Name string `json:"name"`

	// The key where the CA certificate can be found in the Secret or ConfigMap.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MaxLength:=253
	// +kubebuilder:validation:Pattern:=^[-._a-zA-Z0-9]+$
	Key string `json:"key,omitempty"`

	// The namespace the Provider type is in.
	// +optional
	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MaxLength:=63
	// +kubebuilder:validation:Pattern:=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Namespace *string `json:"namespace,omitempty"`
}

type WebhookResult struct {
	// Json path of return value
	// +optional
	JSONPath string `json:"jsonPath,omitempty"`
}

type WebhookSecret struct {
	// Name of this secret in templates
	Name string `json:"name"`

	// Secret ref to fill in credentials
	SecretRef esmeta.SecretKeySelector `json:"secretRef"`
}
