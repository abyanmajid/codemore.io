# codemore.io v0.1.0 (ARCHIVED)

This monorepo contains version 0.1.0 of **codemore.io**, which was my first attempt at building a cloud-native microservices architecture that in theory is scalable and maintainable. This monorepo is **DISCONTINUED DUE TO ABUNDANT INCOMPLIANCES WITH IMPORTANT DEVOPS PRINCIPLES AND BEST PRACTICES**. Moving forward, I will attempt to build **codemore.io** for the second time, taking into account the lessons I have learned from building this monorepo. 

<p align="center">The repositories for <b>codemore.io v0.2.0</b> can be found at <b><a href="https://github.com/YanSystems">@YanSystems</a></b></p>

<!-- <h3 align="center"> Live App 🚀 | Demo 📹 | Documentation 🔍 | Source 📦 </h3> -->

## The old version 0.1.0 architecture

In codemore.io v0.1.0, All client requests are sent to the `broker` service (which serves as an API gateway) with a JSON payload. The `broker` service will then redirect this request via `gRPC` to the correct microservice. The following `mermaid` visualizer depicts the architecture:

```mermaid
graph TD
    Client["<b>Client</b>"]
    Feed["<b>Feed</b><br>(WebSocket)"]
    Broker["<b>Broker</b><br>(REST)"]
    Auth["Third-Party Auth"]
    S3["Cloud Storage"]
    Epsilon["Epsilon CMS"]
    Account["<b>Account</b><br>(REST)"]
    Course["<b>Course</b><br>(REST)"]
    Progression["<b>Progression</b><br>(REST)"]
    Compiler["<b>Compiler</b><br>(gRPC)"]
    Judge["<b>Judge</b><br>(REST, Goroutined gRPC)"]
    CF["<b>Content Fetcher</b><br>(gRPC)"]
    Mail["<b>Mail</b><br>(gRPC)"]
    MongoDB["<b>MongoDB</b>"]

    Client <--> Feed
    Client <--> Auth
    Client <--> S3
    Client <--> Broker

    Broker <--> CF
    Broker <--> Compiler
    Broker <--> Judge
    Broker <--> Account
    Broker <--> Course
    Broker <--> Progression
    Broker <--> Mail

    Account --> MongoDB
    Course --> MongoDB
    Progression --> MongoDB
    Judge <--> Compiler
    Judge <--> MongoDB
    CF <--> Epsilon
    Epsilon --> MongoDB
```
