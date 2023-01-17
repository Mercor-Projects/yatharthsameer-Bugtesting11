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

package io.fabric8.kubernetes.client.mock;

import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.ServiceBuilder;
import io.fabric8.kubernetes.api.model.ServiceList;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.server.mock.KubernetesServer;
import org.junit.Rule;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.migrationsupport.rules.EnableRuleMigrationSupport;

import java.util.Collections;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@EnableRuleMigrationSupport
public class ServiceCrudTest {
  @Rule
  public KubernetesServer server = new KubernetesServer(true, true);

  @Test
  public void testCrud() {
    KubernetesClient client = server.getClient();

    Service service1 = new ServiceBuilder().withNewMetadata().withName("svc1").and().withNewSpec().and().build();
    Service service2 = new ServiceBuilder().withNewMetadata().withName("svc2").addToLabels("foo", "bar").and().withNewSpec().and().build();
    Service service3 = new ServiceBuilder().withNewMetadata().withName("svc3").addToLabels("foo", "bar").and().withNewSpec().and().build();

    client.services().inNamespace("ns1").create(service1);
    client.services().inNamespace("ns2").create(service2);
    client.services().inNamespace("ns1").create(service3);

    ServiceList aServiceList = client.services().list();
    assertNotNull(aServiceList);
    assertEquals(0, aServiceList.getItems().size());

    aServiceList = client.services().inAnyNamespace().list();
    assertNotNull(aServiceList);
    assertEquals(3, aServiceList.getItems().size());

    aServiceList = client.services().inNamespace("ns1").list();
    assertNotNull(aServiceList);
    assertEquals(2, aServiceList.getItems().size());

    client.services().inNamespace("ns1").withName("svc1").delete();
    aServiceList = client.services().inNamespace("ns1").list();
    assertNotNull(aServiceList);
    assertEquals(1, aServiceList.getItems().size());

    aServiceList = client.services().inAnyNamespace().withLabels(Collections.singletonMap("foo", "bar")).list();
    assertNotNull(aServiceList);
    assertEquals(2, aServiceList.getItems().size());

    service2 = client.services().inNamespace("ns2").withName("svc2").edit(s -> new ServiceBuilder(s)
                     .editMetadata().addToLabels("key1", "value1").endMetadata().build());

    assertNotNull(service2);
    assertEquals("value1", service2.getMetadata().getLabels().get("key1"));
  }
}
