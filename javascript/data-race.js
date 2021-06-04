let number = 0;

const sleep = async () => {
  return new Promise((resolve) => setTimeout(resolve, Math.random() * 5));
};

const incrementJob = async () => {
  for (let i = 0; i < 100; i++) {
    await sleep();
    let read = number;
    read = read + 1;
    await sleep();
    number = read;
  }
};

const testDataRace = async () => {
  await Promise.all([incrementJob(), incrementJob()]);
  if (number === 200) {
    throw new Error("Test failed");
  }
  return number;
};

testDataRace()
  .then((x) => console.log(`Test passed: x id ${number}`))
  .catch((err) => console.error(err));
