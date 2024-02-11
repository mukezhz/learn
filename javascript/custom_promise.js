class CustomPromise {
  constructor(executor) {
    this.state = "pending";
    this.value = undefined;
    this.reason = undefined;
    this.onFulfilledCallbacks = [];
    this.onRejectedCallbacks = [];

    const resolve = (value) => {
      if (this.state === "pending") {
        this.state = "fulfilled";
        this.value = value;
        this.onFulfilledCallbacks.forEach((callback) => {
          callback();
        });
      }
    };

    const reject = (reason) => {
      if (this.state === "pending") {
        this.state = "rejected";
        this.reason = reason;
        this.onRejectedCallbacks.forEach((callback) => {
          callback();
        });
      }
    };

    try {
      executor(resolve, reject);
    } catch (error) {
      reject(error);
    }
  }

  then(onFulfilled, onRejected) {
    onFulfilled =
      typeof onFulfilled === "function" ? onFulfilled : (value) => value;
    onRejected =
      typeof onRejected === "function"
        ? onRejected
        : (reason) => {
            throw reason;
          };

    return new CustomPromise((resolve, reject) => {
      const fulfilledTask = () => {
        queueMicrotask(() => {
          try {
            const result = onFulfilled(this.value);
            resolve(result);
          } catch (error) {
            reject(error);
          }
        });
      };

      const rejectedTask = () => {
        queueMicrotask(() => {
          try {
            const result = onRejected(this.reason);
            resolve(result);
          } catch (error) {
            reject(error);
          }
        });
      };

      if (this.state === "fulfilled") {
        fulfilledTask();
      } else if (this.state === "rejected") {
        rejectedTask();
      } else {
        this.onFulfilledCallbacks.push(fulfilledTask);
        this.onRejectedCallbacks.push(rejectedTask);
      }
    });
  }

  catch(onRejected) {
    onRejected =
      typeof onRejected === "function"
        ? onRejected
        : (reason) => {
            throw reason;
          };
    return this.then(undefined, onRejected);
  }
}

function MockFetchCustomPromise(url) {
  return new CustomPromise((resolve, reject) => {
    if (!url) {
      reject("URL is required");
    } else {
      resolve("Data fetched successfully");
    }
  });
}

// Usage
MockFetchCustomPromise("https://example.com/data")
  .then((data) => {
    console.log("Data fetched successfully:", data);
  })
  .catch((error) => {
    console.error("Error fetching data:", error);
  });
