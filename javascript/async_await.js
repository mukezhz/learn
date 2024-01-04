function setTimeoutPromise(delay) {
  return new Promise((resolve, reject) => {
    if (delay < 0) return reject("Delay must be greater than 0")

    setTimeout(() => {
      resolve(`You waited ${delay} milliseconds`)
    }, delay)
  })
}

setTimeoutPromise(250)
  .then(msg => {
    console.log(msg)
    console.log("First Timeout")
    return setTimeoutPromise(500)
  })
  .then(msg => {
    console.log(msg)
    console.log("Second Timeout")
  })

async function doStuff() {
  const msg1 = await setTimeoutPromise(250)
  console.log(msg1)
  console.log("Hello world")
  console.log("First Timeout")

  const msg2 = await setTimeoutPromise(500)
  console.log(msg2)
  console.log("Second Timeout")
  console.log("Hello world")
}
doStuff()