
package io.fabric8.camelk.v1;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.ContainerPort;
import io.fabric8.kubernetes.api.model.EnvVar;
import io.fabric8.kubernetes.api.model.IntOrString;
import io.fabric8.kubernetes.api.model.KubernetesResource;
import io.fabric8.kubernetes.api.model.LabelSelector;
import io.fabric8.kubernetes.api.model.LocalObjectReference;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.api.model.ObjectReference;
import io.fabric8.kubernetes.api.model.PersistentVolumeClaim;
import io.fabric8.kubernetes.api.model.PodTemplateSpec;
import io.fabric8.kubernetes.api.model.ResourceRequirements;
import io.fabric8.kubernetes.api.model.Volume;
import io.fabric8.kubernetes.api.model.VolumeMount;
import io.sundr.builder.annotations.Buildable;
import io.sundr.builder.annotations.BuildableReference;
import lombok.EqualsAndHashCode;
import lombok.Setter;
import lombok.ToString;
import lombok.experimental.Accessors;

@JsonDeserialize(using = com.fasterxml.jackson.databind.JsonDeserializer.None.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({
    "apiVersion",
    "kind",
    "metadata",
    "artifactId",
    "dependencies",
    "groupId",
    "languages",
    "version"
})
@ToString
@EqualsAndHashCode
@Setter
@Accessors(prefix = {
    "_",
    ""
})
@Buildable(editableEnabled = false, validationEnabled = false, generateBuilderPackage = false, lazyCollectionInitEnabled = false, builderPackage = "io.fabric8.kubernetes.api.builder", refs = {
    @BuildableReference(ObjectMeta.class),
    @BuildableReference(LabelSelector.class),
    @BuildableReference(Container.class),
    @BuildableReference(PodTemplateSpec.class),
    @BuildableReference(ResourceRequirements.class),
    @BuildableReference(IntOrString.class),
    @BuildableReference(ObjectReference.class),
    @BuildableReference(LocalObjectReference.class),
    @BuildableReference(PersistentVolumeClaim.class),
    @BuildableReference(EnvVar.class),
    @BuildableReference(ContainerPort.class),
    @BuildableReference(Volume.class),
    @BuildableReference(VolumeMount.class)
})
public class CamelLoader implements KubernetesResource
{

    @JsonProperty("artifactId")
    private java.lang.String artifactId;
    @JsonProperty("dependencies")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private List<MavenArtifact> dependencies = new ArrayList<MavenArtifact>();
    @JsonProperty("groupId")
    private java.lang.String groupId;
    @JsonProperty("languages")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private List<java.lang.String> languages = new ArrayList<java.lang.String>();
    @JsonProperty("metadata")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private Map<String, String> metadata = new LinkedHashMap<String, String>();
    @JsonProperty("version")
    private java.lang.String version;
    @JsonIgnore
    private Map<java.lang.String, Object> additionalProperties = new HashMap<java.lang.String, Object>();

    /**
     * No args constructor for use in serialization
     * 
     */
    public CamelLoader() {
    }

    /**
     * 
     * @param metadata
     * @param languages
     * @param groupId
     * @param artifactId
     * @param version
     * @param dependencies
     */
    public CamelLoader(java.lang.String artifactId, List<MavenArtifact> dependencies, java.lang.String groupId, List<java.lang.String> languages, Map<String, String> metadata, java.lang.String version) {
        super();
        this.artifactId = artifactId;
        this.dependencies = dependencies;
        this.groupId = groupId;
        this.languages = languages;
        this.metadata = metadata;
        this.version = version;
    }

    @JsonProperty("artifactId")
    public java.lang.String getArtifactId() {
        return artifactId;
    }

    @JsonProperty("artifactId")
    public void setArtifactId(java.lang.String artifactId) {
        this.artifactId = artifactId;
    }

    @JsonProperty("dependencies")
    public List<MavenArtifact> getDependencies() {
        return dependencies;
    }

    @JsonProperty("dependencies")
    public void setDependencies(List<MavenArtifact> dependencies) {
        this.dependencies = dependencies;
    }

    @JsonProperty("groupId")
    public java.lang.String getGroupId() {
        return groupId;
    }

    @JsonProperty("groupId")
    public void setGroupId(java.lang.String groupId) {
        this.groupId = groupId;
    }

    @JsonProperty("languages")
    public List<java.lang.String> getLanguages() {
        return languages;
    }

    @JsonProperty("languages")
    public void setLanguages(List<java.lang.String> languages) {
        this.languages = languages;
    }

    @JsonProperty("metadata")
    public Map<String, String> getMetadata() {
        return metadata;
    }

    @JsonProperty("metadata")
    public void setMetadata(Map<String, String> metadata) {
        this.metadata = metadata;
    }

    @JsonProperty("version")
    public java.lang.String getVersion() {
        return version;
    }

    @JsonProperty("version")
    public void setVersion(java.lang.String version) {
        this.version = version;
    }

    @JsonAnyGetter
    public Map<java.lang.String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    @JsonAnySetter
    public void setAdditionalProperty(java.lang.String name, Object value) {
        this.additionalProperties.put(name, value);
    }

}