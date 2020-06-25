import { Node } from 'flatend';

let counter = 0;

const count = (ctx: any) => ctx.send(`${counter++}`);

const main = async () => {
  await Node.start({
    addrs: ['127.0.0.1:9000'],
    services: { count }
  });
};

main().catch((err) => console.error(err));
