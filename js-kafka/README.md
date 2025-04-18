
---

## 🧱 Core Kafka Concepts to Implement

### 1. **Broker (Server)**
The central server that stores messages and serves client requests (producers/consumers).

**Responsibilities:**
- Receive messages from producers
- Store messages in a persistent, append-only log
- Serve messages to consumers

### 2. **Producer**
A client that sends data (messages) to the broker.

**Responsibilities:**
- Connect to the broker
- Choose topic and partition (if applicable)
- Serialize and send messages

### 3. **Consumer**
A client that reads data from the broker.

**Responsibilities:**
- Connect to broker
- Subscribe to topic(s)
- Track offsets (which message has been read)
- Pull messages from the broker

### 4. **Topic & Partition**
Kafka topics are split into partitions (think logs). Each partition is an ordered, immutable sequence of messages.

**Key Idea:**
- Partitioning allows parallelism and scalability.
- Each partition is a commit log.

### 5. **Offset Tracking**
Offsets are positions of messages in a partition. Consumers use offsets to know where they are in the log.

---