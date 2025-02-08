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
    new sst.aws.Astro("TodoPlannerAstroWebsite", {
      domain: {
        name: "todoplanner.app",
        dns: sst.cloudflare.dns({
          zone: "2a190c1e0fb8d59ad02cb4646c8b6ce9",
        }),
      },
      path: "./packages/www",
      buildCommand: "bun run build",
      dev: {
        command: "bun run dev",
      },
    });

    /**
     * Setting up a Virtual Private Cloud (VPC) to establish
     * a secure and isolated connection for containerized applications.
     *
     * This ensure scalability, security and efficent resource management
     * for workloads running in the container.
     */
    const todoPlannerVpc = new sst.aws.Vpc("TodoPlannerVPC", {});

    /**
     * Creating an Amazon ECS Cluster to manage and orchestrate
     * containerized applications within the specified VPC.
     */
    const ecsCluster = new sst.aws.Cluster("TodoPlannerCluster", {
      vpc: todoPlannerVpc,
      forceUpgrade: "v2",
    });
  },
});
