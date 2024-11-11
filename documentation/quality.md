# Quality Models
Below are quality models for Software, Data and Models. They serve as 
metrics to measure this code against throughout the development lifecycle.

## Software QM
Adapted from ISO/IEC 25010:2011

### Functional Suitability
Does the software meet the needs for specified tasks under specified conditions?

| Functional Suitability         | Definition                                                                                        |
|--------------------------------|---------------------------------------------------------------------------------------------------|
| **Functional Completeness**    | The extent to which the software provides all the necessary functionality as required by users.   |
| **Functional Correctness**     | The accuracy and correctness with which the software performs its intended functions.             |
| **Functional Appropriateness** | The suitability of the functions provided by the software for the specified tasks and objectives. |

### Performance Efficiency
How well does the software use resources relative to the performance it delivers?

| Performance Efficiency   | Definition                                                                                               |
|--------------------------|----------------------------------------------------------------------------------------------------------|
| **Time Behavior**        | The responsiveness and processing time of the software under specified conditions.                       |
| **Resource Utilization** | The efficiency with which the software uses system resources (CPU, memory, disk, etc.) to perform tasks. |
| **Capacity**             | The degree to which the software can handle a defined amount of data or requests over time.              |
| **Scalability**          | The extent to which the software can handle an increasing load by scaling up or out.                     |

### Compatibility
How well does the software operate alongside other systems?

| Compatibility        | Definition                                                                                                        |
|----------------------|-------------------------------------------------------------------------------------------------------------------|
| **Co-existence**     | The ability of the software to operate efficiently while sharing resources with other software.                   |
| **Interoperability** | The capability of the software to interact with other systems or components and exchange information effectively. |

### Usability
How easy is it for users to understand, learn, and operate the software?

| Usability                           | Definition                                                                               |
|-------------------------------------|------------------------------------------------------------------------------------------|
| **Appropriateness Recognizability** | The degree to which users can identify if the software is suitable for their needs.      |
| **Learnability**                    | The ease with which users can learn to use the software and understand its features.     |
| **Operability**                     | The ability of the software to allow users to control and operate it effectively.        |
| **User Error Protection**           | The extent to which the software prevents or minimizes user errors.                      |
| **User Interface Aesthetics**       | The appeal and visual clarity of the software’s interface design.                        |
| **Accessibility**                   | The degree to which the software can be used by people with disabilities or limitations. |

### Reliability
How well does the software function under normal and challenging conditions?

| Reliability         | Definition                                                                                 |
|---------------------|--------------------------------------------------------------------------------------------|
| **Maturity**        | The ability of the software to avoid failure as a result of defects.                       |
| **Availability**    | The probability that the software will be operational and accessible when needed.          |
| **Fault Tolerance** | The extent to which the software can maintain performance in the event of a failure.       |
| **Recoverability**  | The capability of the software to recover its functions in the event of a failure.         |
| **Resilience**      | The ability of the software to maintain or recover functionality under failure conditions. |

### Security
How well does the software protect information and resist unauthorized access?

| Security            | Definition                                                                                       |
|---------------------|--------------------------------------------------------------------------------------------------|
| **Confidentiality** | The degree to which the software prevents unauthorized access to data and information.           |
| **Integrity**       | The capability of the software to prevent unauthorized modifications of data.                    |
| **Non-repudiation** | The ability to prove the actions and events performed by users to prevent denial of involvement. |
| **Accountability**  | The ability to trace actions and identify responsible users or components.                       |
| **Authenticity**    | The assurance that data and entities involved in a communication are genuine.                    |
| **Compliance**      | The adherence to regulatory standards and security or privacy guidelines.                        |

### Maintainability
How well can the software be modified for updates or changes?

| Maintainability   | Definition                                                                                                        |
|-------------------|-------------------------------------------------------------------------------------------------------------------|
| **Modularity**    | The extent to which the software is composed of discrete components that can be modified independently.           |
| **Reusability**   | The ability of parts of the software to be used in other projects or applications.                                |
| **Analyzability** | The ease with which the software can be analyzed for issues and areas of improvement.                             |
| **Modifiability** | The extent to which the software can be adapted to changes without introducing defects.                           |
| **Testability**   | The ease with which the software can be tested to verify that it performs as expected.                            |
| **Observability** | The capability to monitor, trace, and measure the software’s performance and behavior in production environments. |

### Portability
How easily can the software be transferred to other environments?

| Portability        | Definition                                                                                    |
|--------------------|-----------------------------------------------------------------------------------------------|
| **Adaptability**   | The ease with which the software can be adapted to different environments or configurations.  |
| **Installability** | The ease with which the software can be successfully installed in a target environment.       |
| **Replaceability** | The ability of the software to replace or be replaced by another, achieving the same purpose. |

## Data QM
Based on ISO 25012

### Intrinsic Quality
How well does the data represent actual, true information?

| Intrinsic Quality     | Definition                                                                                                          |
|-----------------------|---------------------------------------------------------------------------------------------------------------------|
| **Accuracy**          | The extent to which data correctly represents the real-world values or conditions it is meant to model.             |
| **Consistency**       | The degree to which data remains consistent across different datasets, sources, and times.                          |
| **Completeness**      | The degree to which all required data values are present and recorded.                                              |
| **Validity**          | The extent to which data conforms to the required formats, rules, or constraints.                                   |
| **Uniqueness**        | The extent to which data records or elements are not duplicated within the dataset.                                 |
| **Bias and Fairness** | The extent to which data avoids biases or disparities, especially when used for decision-making and model training. |
| **Traceability**      | The ability to trace data back to its origin to verify authenticity and understand data lineage.                    |

### Contextual Quality
How relevant and meaningful is the data in context?

| Contextual Quality            | Definition                                                                                              |
|-------------------------------|---------------------------------------------------------------------------------------------------------|
| **Relevance**                 | The extent to which data is applicable to the specific analysis, business, or use case requirements.    |
| **Timeliness**                | The extent to which data is available when needed, relative to the demands of the use case.             |
| **Completeness (in context)** | The extent to which all required data is present for a specific analysis or use case.                   |
| **Appropriateness**           | The extent to which data’s detail level or format suits the intended analysis or application.           |
| **Ethical Compliance**        | The adherence to data privacy, protection, and ethical standards, especially in sensitive applications. |

### Representational Quality
How well does the data structure and presentation facilitate use?

| Representational Quality            | Definition                                                                                       |
|-------------------------------------|--------------------------------------------------------------------------------------------------|
| **Interpretability**                | The ease with which users can understand data content and context.                               |
| **Consistency (of representation)** | The uniformity of data format, labels, or presentation across different sources and uses.        |
| **Accessibility**                   | The extent to which data is available and retrievable when needed by authorized users.           |
| **Format Compliance**               | The degree to which data adheres to specific, predefined formats or standards.                   |
| **Traceability**                    | The ability to trace data back to its origin to verify authenticity and understand data lineage. |

### Accessibility and Availability
How easily can data be accessed, and how reliable is access?

| Accessibility and Availability | Definition                                                                                 |
|--------------------------------|--------------------------------------------------------------------------------------------|
| **Availability**               | The degree to which data is present and accessible when required.                          |
| **Security/Access Control**    | The extent to which data access is protected and available only to authorized individuals. |
| **Reliability**                | The extent to which data access is consistently maintained and unimpeded.                  |
| **Recovery**                   | The ability to restore data access after a disruption or failure.                          |

### Operational Quality
How well does the data system perform to meet quality standards?

| Operational Quality              | Definition                                                                                                     |
|----------------------------------|----------------------------------------------------------------------------------------------------------------|
| **Data Processing Efficiency**   | The ability to process and retrieve data without unnecessary delays or resource wastage.                       |
| **Scalability**                  | The ability of data systems to handle an increasing amount of data without loss of quality or efficiency.      |
| **Monitoring and Observability** | The extent to which data quality can be effectively monitored in real-time or during processing cycles.        |
| **Maintainability**              | The ease with which data and metadata structures can be maintained over time.                                  |
| **Auditability**                 | The ease with which data and transformations can be reviewed and verified for compliance or accuracy purposes. |

## Machine Learning QM
Adapted from the Booking.com ML Quality Model

### Utility
Does the system provide functions that meet needs under specified conditions?

| Utility            | Definition                                                                                         |
|--------------------|----------------------------------------------------------------------------------------------------|
| **Accuracy**       | The size of the observational error (systematic and random) of an ML system.                       |
| **Effectiveness**  | The ability of an ML system to produce a desired result on the business task it is being used for. |
| **Responsiveness** | The ability of an ML system to complete the desired task in an acceptable time frame.              |
| **Reusability**    | The ability to reuse the same ML system without any change, for another business case.             |

### Economy
Does it perform well given resources and conditions?

| Economy                | Definition                                                                                                                                        |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| **Cost-effectiveness** | The extent to which an ML System achieves the desired relationship between costs and overall impact.                                              |
| **Efficiency**         | The ability to avoid wasting resources (computational, human, financial, etc.) in order to perform the desired task effectively.                  |
| **Availability**       | The probability that the system will be up and running and able to deliver useful services to users.                                              |
| **Resilience**         | The extent to which an ML system can provide and maintain an acceptable level of service in the face of technical challenges to normal operation. |
| **Adaptability**       | The extent to which an ML system can adapt to changes in the production environment, always providing the same functioning level.                 |
| **Scalability**        | The extent of an ML system to handle a growing amount of work by adding resources to the system.                                                  |

### Robustness
How little does the system degrade under adverse, changing, or unexpected conditions?

| Robustness          | Definition                                                                                                             |
|---------------------|------------------------------------------------------------------------------------------------------------------------|
| **Stability**       | The extent to which an ML system can maintain performance in the face of load, environment, or infrastructure changes. |
| **Fault Tolerance** | The ability of an ML system to continue operating properly in the event of the failure of some of its components.      |
| **Recoverability**  | The ease with which an ML system can be restored to normal functioning after a disruption.                             |

### Modifiability
How well is it suited for adaptations?

| Modifiability       | Definition                                                                                                         |
|---------------------|--------------------------------------------------------------------------------------------------------------------|
| **Extensibility**   | The ease with which a system can be modified, in order to be used for another purpose.                             |
| **Maintainability** | The ease with which activities aiming at keeping an ML system functional in the desired regime, can be performed.  |
| **Modularity**      | The extent to which an ML system consists of components of limited functionality that interrelate with each other. |
| **Testability**     | The extent with which an ML system can be tested against expected behaviors.                                       |

### Productionizability
How well can it execute in production?

| Productionizability | Definition                                                                                                                                |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| **Deployability**   | The extent to which an ML system can be deployed in production when needed.                                                               |
| **Repeatability**   | The ease with which the ML lifecycle can be repeated.                                                                                     |
| **Operability**     | The extent to which an ML system can be controlled (disable, revert, upload new version, etc.).                                           |
| **Monitoring**      | The extent to which relevant indicators of an ML system are effectively observed/monitored and integrated in the operation of the system. |

### Comprehensibility
How easy is it to understand/onboard?

| Comprehensibility     | Definition                                                                                                                                                                                                                                                                                                       |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Discoverability**   | The extent to which an ML system can be found (by means of performing a search in a database or any other information retrieval system).                                                                                                                                                                         |
| **Readability**       | The ease with which the code of an ML system can be understood.                                                                                                                                                                                                                                                  |
| **Traceability**      | The ability to relate artifacts created during the development of an ML system to describe the system from different perspectives and levels of abstraction with each other, the stakeholders that have contributed to the creation of the artifacts, and the rationale that explains the form of the artifacts. |
| **Understandability** | The ease with which the implementation and design choices of an ML system can be understood.                                                                                                                                                                                                                     |
| **Usability**         | The extent to which an ML system can be effectively used by users.                                                                                                                                                                                                                                               |
| **Debuggability**     | The extent to which the inner workings of the ML system can be analyzed in order to understand why it behaves the way it does.                                                                                                                                                                                   |

### Responsibility
How trustworthy is it?

| Responsibility           | Definition                                                                                                                   |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------|
| **Explainability**       | The ability to explain the output of an ML system.                                                                           |
| **Fairness**             | The extent to which an ML system prevents unjust predictions towards protected attributes (race, gender, income, etc).       |
| **Ownership**            | The extent to which there exists people appointed to maintaining the ML System and supporting all the relevant stakeholders. |
| **Standards Compliance** | The extent to which applicable standards are followed in the ML system.                                                      |
| **Vulnerability**        | The ease with which the system can be (ab)used to achieve malicious purposes.                                                |

