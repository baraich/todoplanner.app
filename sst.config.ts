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
        tls: "5.1.0",
      },
    };
  },
  async run() {
    /**
     * Deploying the application to cloudflare pages.
     * Makes sure that the website is up and running almost cost free,
     * with DDoS protection out of the box.
     */
    new sst.cloudflare.StaticSite("TodoPlannner", {
      path: "./packages/www",
      build: {
        command: "bun run build",
        output: "./dist",
      },
      domain: "todoplanner.app",
    });

    /**
     * Wrapping the TLS.PrivateKey using a Linkable allowing for
     * easy integration with other services.
     */
    sst.Linkable.wrap(tls.PrivateKey, (resource) => ({
      properties: {
        public: resource.publicKeyOpenssh,
        private: resource.privateKeyOpenssh,
      },
    }));

    /**
     * Creating a new TLS.PrivateKey resource for the TodoPlanner
     * application using the ED25519 algorithm, allowing secure connections
     * from clients over SSH.
     */
    const todoPlannerKey = new tls.PrivateKey("TodoPlannerKey", {
      algorithm: "ED25519",
    });

    /**
     * Setting up a Virtual Private Cloud (VPC) to establish
     * a secure and isolated connection for containerized applications.
     *
     * This ensure scalability, security and efficent resource management
     * for workloads running in the container.
     */
    const todoPlannerVpc = new sst.aws.Vpc("TodoPlannerVPC", {
      nat: "ec2",
      bastion: true,
    });

    /**
     * Creating an Amazon ECS Cluster to manage and orchestrate
     * containerized applications within the specified VPC.
     */
    const ecsCluster = new sst.aws.Cluster("TodoPlannerCluster", {
      vpc: todoPlannerVpc,
      forceUpgrade: "v2",
    });

    /**
     * Deploying and ECS service for managing secure SSH access to
     * TodoPlanner application. Service is configured to use minimal
     * resource (0.25 vCPU and 0.5 GB memory) to keep costs low.
     *
     * The service is created behind a load balancer, which listens on
     * 22/tcp and forwards traffic to 2222/tcp. The load balancer is
     * accociated with "todoplanner.app" fomain handled via Cloudflare.
     *
     * Scalling is intentionally explicity limited to a single resource,
     * to keep the running cost minimal.
     */
    ecsCluster.addService("SSH", {
      cpu: "0.25 vCPU",
      memory: "0.5 GB",
      link: [todoPlannerKey],
      loadBalancer: {
        domain: {
          name: "cli.todoplanner.app",
          dns: sst.cloudflare.dns(),
        },
        rules: [{ listen: "22/tcp", forward: "2222/tcp" }],
      },
      scaling: {
        min: 1,
        max: 1,
      },
      image: {
        context: "./packages/app",
      },
    });
  },
});
