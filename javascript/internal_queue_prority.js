process.nextTick(() => console.log("this is process.nextTick 1"));
process.nextTick(() => {
  console.log("this is process.nextTick 2");
  process.nextTick(() =>
    console.log("[BEFORE PROMISE]this is the inner next tick inside next tick")
  );
});
process.nextTick(() => console.log("this is process.nextTick 3"));

Promise.resolve().then(() => console.log("this is Promise.resolve 1"));
Promise.resolve().then(() => {
  console.log("this is Promise.resolve 2");
  process.nextTick(() =>
    console.log("this is the inner next tick inside Promise then block")
  );
});
Promise.resolve().then(() => console.log("this is Promise.resolve 3"));

process.nextTick(() => console.log("this is process.nextTick I"));
process.nextTick(() => {
  console.log("this is process.nextTick II");
  process.nextTick(() =>
    console.log("[AFTER PROMISE]this is the inner next tick inside next tick")
  );
});
process.nextTick(() => console.log("this is process.nextTick III"));
