/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "todoplanner-app",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
      providers: {
        aws: {
          region: "ap-south-1",
        },
      },
    };
  },
  async run() {
    new sst.cloudflare.StaticSite("TodoPlannner", {
      path: "./packages/www",
      build: {
        command: "bun run build",
        output: "./dist",
      },
      domain: "todoplanner.app",
    });

    /**
     * Setting up a Virtual Private Cloud (VPC) to establish
     * a secure and isolated connection for containerized applications.
     *
     * This ensure scalability, security and efficent resource management
     * for workloads running in the container.
     */
    // TODO: Uncomment this when we have a use case for it
    // const todoPlannerVpc = new sst.aws.Vpc("TodoPlannerVPC", {});

    /**
     * Creating an Amazon ECS Cluster to manage and orchestrate
     * containerized applications within the specified VPC.
     */
    // TODO: Uncomment this when we have a use case for it
    // const ecsCluster = new sst.aws.Cluster("TodoPlannerCluster", {
    //   vpc: todoPlannerVpc,
    //   forceUpgrade: "v2",
    // });
  },
});
