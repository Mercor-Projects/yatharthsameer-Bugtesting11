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
package io.fabric8.kubernetes.client.utils.internal;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.atomic.AtomicInteger;

public class ExponentialBackoffIntervalCalculator {

  private static final Logger logger = LoggerFactory.getLogger(ExponentialBackoffIntervalCalculator.class);

  private final int initialInterval;
  private final int maxRetryIntervalExponent;
  final AtomicInteger currentReconnectAttempt = new AtomicInteger();

  public ExponentialBackoffIntervalCalculator(int initialInterval, int maxRetryIntervalExponent) {
    this.initialInterval = initialInterval;
    this.maxRetryIntervalExponent = maxRetryIntervalExponent;
  }

  public long getInterval(int retryIndex) {
    int exponentOfTwo = retryIndex;
    if (exponentOfTwo > maxRetryIntervalExponent) {
      exponentOfTwo = maxRetryIntervalExponent;
    }
    return (long) initialInterval * (1 << exponentOfTwo);
  }

  public void resetReconnectAttempts() {
    currentReconnectAttempt.set(0);
  }

  public final long nextReconnectInterval() {
    int exponentOfTwo = currentReconnectAttempt.getAndIncrement();
    long ret = getInterval(exponentOfTwo);
    logger.debug("Current reconnect backoff is {} milliseconds (T{})", ret, exponentOfTwo);
    return ret;
  }

  public int getCurrentReconnectAttempt() {
    return currentReconnectAttempt.get();
  }

}
