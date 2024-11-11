# Architecture

## Introduction
The yunque backend is heavily tailored to ideas regarding the 
user-experience and marketing strategies. At its core it is a 
dynamically-scalable microservice architecture that deploys easily in 
kubernetes, and lives in a monorepo (mostly because it is a personal project).

In this repository there are both library modules and services, the 
difference being that services are deployed as individual units. These services 
fall under 3 different categories:
- **Engines** are medium to large services that execute business logic and
  handle requests
- **Agents** are small services that either monitor or augment the operations 
  of the other services present.
- **Insights Services** are analytics services, they operate under slightly 
  different scaling requirements and execute the primary function of 
  streamlining data from the other services as well as historical sources 
  into data pipelines. After processing, the data is made actionable in 
  analysis/visualization services and algorithms/models.

## Engines Overview
### Gateway
Almost all requests pass through a central API Gateway, which limits the 
amount of services that need to expose public endpoints. The gateway not 
only reroutes traffic, but executes secondary functions such as batching 
requests, prefetching, intercepting interaction data and sending canaries. 
Because this is such an important service, there is a lighter backup version 
that idles and can be switched out if needed.

### Shop
The original idea for the architecture had separate product, cart, 
reviews and inventory services. These have since been combined into one shop 
service, representing the idea that the store a user is interfacing with is 
simply an extension of their cart, and that items either fall into the 
category of "is buying" or "is likely to buy". This property makes the 
functions of products and the user cart strongly coupled.

### Catalogue
While the shop service has the ability to track and update inventory for 
products, it represents the user-facing side of the catalogue, and therefore 
there is a separate catalogue service that represents a seller-facing side 
of the store. Both services persist to a single distributed source of truth, 
which is managed by helper services.

### Orders
The transition between a cart and checkout is a major shift in the state of 
a customer. Therefore, an independent service exists to manage orders - from 
their inception to payment to delivery. Since the checkout process on a 
user's device is relatively slow compared to other operations, optimizations 
on this transition are left to the frontend.

### Payment
Payments can be in regard to orders, subscriptions or refunds, and thus are 
separated into a standalone engine. The payment service coordinates payments,
but does not provide the logic for payments, which is outsourced to external 
services. Instead, it acts as a central hub for all payments.

### Notifications
Many actions in the chain can generate notifications that need to be sent 
per email, text or over some other medium. A simple notification service 
handles all of these messages.

### Session Service
Operating as an edge-optimized service for persisting and caching data of 
active users, the session service is deployed at edge nodes.

## Agents Overview
### Health Agent
Since this is a more abstract personal project and doesn't have a high 
deployment footprint, I will focus on a small number (starting with one) of
health agents. This small monitoring service will send canary requests and 
monitor the availability and speed of other services.

### Quality Agent
This agent measures quality according to predefined quality models. Having a 
good model of where a system measures according to agreed-upon requirements 
allows us to forecast impacts of queued deployments or autoscaling on 
these metrics in a principled way.

### Security Agent
While this project doesn't dive too deep into security apart from the code 
itself being written in a way that is secure, this agent probes systems and 
searches for vulnerabilities while the service is in production. The 
difference between this and testing stages is that the agent does not 
deliberately poke holes and generate vulnerabilities in in-production 
services - that would be a terrible idea - instead it monitors the state of 
these services and searches for potential vulnerabilities that may arise 
from state changes in production.

## Analytics Overview
The analytics service provides an entryway into analytics for this app. Over 
time this monolith will be replaced by a more deliberate system.