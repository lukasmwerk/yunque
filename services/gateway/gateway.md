# Yunque Gateway

The gateway service serves as an api gateway to the rest of the backend. It is important to note that the gateway service must be available for other internal services to be reachable from other parties. This allows backend services to communicate with more efficient protocols over local networks, while providing a failsafe interface that protects callers from potential failures of internal services. The gateway service is built to be robust and replicable, meaning any location running a backend service can run its own gateway gateway, as a coordinator node.

While the gateway service provides a barrier between the backend and the rest of the world, services can contact eachother bypassing the gateway to prevent bottlenecks and overcomplicated flows.

The goal of the gateway service is to be lightweight enough to be easily spun up by any underlying orchestration.