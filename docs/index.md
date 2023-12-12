---
hide:
- navigation
- toc
---

![logo](imgs/homelogo.png)

<br/>
<br/>

## 📣 tfsecurity to Tunnel Migration

As part of our goal to provide a comprehensive open source security solution for all, we have been consolidating all of our scanning-related efforts in one place, and that is [Tunnel](https://github.com/khulnasoft/tunnel). 

Over the past year, tfsecurity has laid the foundations to Tunnel's IaC & misconfigurations scanning capabilities, including Terraform scanning, which has been natively supported in Tunnel for a long time now.

Going forward we want to encourage the tfsecurity community to transition over to Tunnel. Moving to Tunnel gives you the same excellent Terraform scanning engine, with some extra benefits:

1. Access to more languages and features in the same tool.
2. Access to more integrations with tools and services through the rich ecosystem around Tunnel.
3. Commercially supported by KhulnaSoft as well as by a the passionate Tunnel community.
tfsecurity will continue to remain available for the time being, although our engineering attention will be directed at Tunnel going forward.

## tfsecurity to Tunnel migration guide

For further information on how Tunnel compares to tfsecurity and moving from tfsecurity to Tunnel, do have a look at the [migration guide.](https://github.com/khulnasoft-lab/tfsecurity/blob/master/tfsecurity-to-vul-migration-guide.md)

## tfsecurity

`tfsecurity` is a static analysis security scanner for your Terraform code.

Designed to run locally and in your CI pipelines, developer-friendly output and fully documented checks mean detection and remediation can take place as quickly and efficiently as possible

`tfsecurity` takes a developer-first approach to scanning your Terraform templates; using static analysis and deep integration with the official HCL parser it ensures that security issues can be detected before your infrastructure changes take effect.

<br/>
<br/>


<figure style="text-align: center">
  <img src="imgs/demo.gif" width="1000">
  <figcaption>Demo: Misconfiguration Detection</figcaption>
</figure>

`tfsecurity` is an [KhulnaSoft Security][khulnasoft] open source project.  
Learn about our open source work and portfolio [here][oss].  
Contact us about any matter by opening a GitHub Discussion [here][discussions]


[khulnasoft]: https://khulnasoft.com
[oss]: https://www.khulnasoft.com/products/open-source-projects/
[discussions]: https://github.com/khulnasoft-lab/tfsecurity/discussions
