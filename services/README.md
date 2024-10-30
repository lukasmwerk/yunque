# Services
Microservices for the yunque backend

This architecture does not prevent upstream errors in services from causing failures in the core system, but it is designed to be fairly resistant to these failures.

This README serves as an entry point for (micro)service documentation.

## [analytics](analytics)
A service for processing analytics both for internal dynamic updates and admin decision making

## [auth](auth)
A simple authorization service for logging in to the store

# [core](core)
The main backend api-gateway that connects the microservices to end-users

## [help](help)
A centralized help platform for providing help services, customer service, etc

## [inventory](inventory)
An inventory management service

## [notification](notification)
A simple notification service to send email updates on orders, subscriptions

## [orders](orders)
Yunque's order management system

## [payment](payment)
A simple payment system

## [product](product)
A product management microservice that handles CRUD operations on products and integrates with analytics

## [reviews](reviews)
A review system for customers and sellers to leave reviews