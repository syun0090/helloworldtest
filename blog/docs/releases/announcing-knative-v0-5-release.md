---
title: "v0.5 release"
linkTitle: "v0.5 release"
date: 2019-04-11
description: "The Knative v0.5 release announcement"
type: "blog"
image: knative-eventing.png
---

<body>
<article class="h-entry">
<section data-field="subtitle" class="p-summary">
Once again, we are excited to announce new release of Knative: a
platform to help developers build, deploy, and manage modern serverless…
</section>
<section data-field="body" class="e-content">
<section name="3f42" class="section section--body section--first section--last" >
<div class="section-divider"><hr class="section-divider" /></div>
<div class="section-content">
<div class="section-inner sectionLayout--insetColumn">
<h3 name="da2e" id="da2e" class="graf graf--h3 graf--leading graf--title" >
Announcing Knative v0.5 Release
</h3>
<p name="bc75" id="bc75" class="graf graf--p graf-after--h3">
Once again, we are excited to announce a new release of
<a href="https://www.knative.dev/" data-href="https://www.knative.dev/" class="markup--anchor markup--p-anchor" rel="noopener" target="_blank" >Knative</a >: a platform to help developers build, deploy, and manage modern serverless workloads on Kubernetes. </p>
<p name="dd43" id="dd43" class="graf graf--p graf-after--p"> While the more frequent and predictable releases of Knative gives us an opportunity to collect faster feedback from real-world use-cases, they also mean smaller and more incremental features. Well, that’s not always the case. Knative v0.5 delivers an exciting set of updates in eventing. Introducing Trigger and Broker objects which further improve and simplify the developer experience building event-driven systems on Knative. </p>
<p name="3634" id="3634" class="graf graf--p graf-after--p"> In addition to eventing, this release of Knative also improves a number of metrics and the overall observability of autoscaling, queue proxy, and Istio telemetry. Let’s review these and a few other changes in more depth: </p>
<h3 name="e284" id="e284" class="graf graf--h3 graf-after--p"> Eventing </h3>
<p name="b67b" id="b67b" class="graf graf--p graf-after--h3"> With the introduction of <strong class="markup--strong markup--p-strong">Trigger</strong> and <strong class="markup--strong markup--p-strong">Broker</strong> objects into the <strong class="markup--strong markup--p-strong" >Eventing</strong > architecture, developers can easily build robust, complex, event-driven systems. By decoupling <strong class="markup--strong markup--p-strong" >Producing</strong > and <strong class="markup--strong markup--p-strong" >Consuming</strong > services there is no longer the need for complex wiring or routing configuration. We are excited to see what new types of events and innovative solutions the community will develop using this new capability! </p>
<figure name="b2df" id="b2df" class="graf graf--figure graf-after--p" >
<div class="aspectRatioPlaceholder is-locked" style="max-width: 700px; max-height: 291px;" >
<div class="aspectRatioPlaceholder-fill" ></div>
<img class="graf-image" data-image-id="knative-eventing.png" data-width="2586" data-height="1076" data-is-featured="true" src="/blog/images/knative-eventing.png" />
</div>
<figcaption class="imageCaption">
Knative Eventing Object Model
</figcaption>
</figure>
<p name="dafe" id="dafe" class="graf graf--p graf-after--figure">
<strong class="markup--strong markup--p-strong" >Trigger: </strong >developers no longer need to manually provision transport for their events and route them to downstream knative services. They simply define an event trigger that selects the source events (with any desired filtering) and sends them to the consuming service. This greatly simplifies the developer experience. </p>
<p name="ad1c" id="ad1c" class="graf graf--p graf-after--p"> <strong class="markup--strong markup--p-strong">Broker: </strong >the events Broker serves as the central events hub to which all messages are sent. Developers and Users simply write services or configure Sources that emit events to the Broker, which handles the rest. Consuming services need only create Triggers to receive the events in which they’re interested in from the Broker. </p>
<p name="cfe0" id="cfe0" class="graf graf--p graf-after--p"> <strong class="markup--strong markup--p-strong" >New event Source: </strong >this release of Knative adds support for the Kafka event source, which brings the power and richness of the Kafka ecosystem to Knative and Kubernetes. </p>

<h3 name="1397" id="1397" class="graf graf--h3 graf-after--p"> Autoscaling </h3> <p name="112e" id="112e" class="graf graf--p graf-after--h3"> Autoscaling added improvements which makes autoscaling under a variety of workloads a smoother motion as well as being more efficient. Expansion of autoscaling metrics were added for additional visibility over time-frames. </p>

<h3 name="65a6" id="65a6" class="graf graf--h3 graf-after--p"> Core API </h3> <p name="dbe8" id="dbe8" class="graf graf--p graf-after--h3"> In this release, named sub-routes now surface their URLs in the status of Service and Route resources, so there’s no more guesswork in how to target one fork of your traffic split. This is one of the first changes to result from our “v1beta1 task force” which has been discussing the next iteration of the Serving API. Expect to see lots more changes in the coming releases. </p>

<p name="48d8" id="48d8" class="graf graf--p graf-after--p"> In addition, several of the default values populated by our webhook are now configurable through a new ConfigMap called config-defaults. We have also increased the visibility into system errors by surfacing more Kubernetes events when our controllers suffer internal errors. Last but not least, we have expanded our conformance testing to include securityContext and metadata.generateName. </p>

<h3 name="f9b0" id="f9b0" class="graf graf--h3 graf-after--p"> Networking </h3> <p name="7e4b" id="7e4b" class="graf graf--p graf-after--h3"> Bulk of the work in the networking space this sprint focused on fixing bugs and improving overall cold-starts for gRPC services as well as further improving client default authority header handling. </p>

<p name="3590" id="3590" class="graf graf--p graf-after--p"> The complete set of Knative v0.5 release notes outlining the new features as well as bug fixes are available in the <a href="https://github.com/knative/serving/releases/tag/v0.5.0" data-href="https://github.com/knative/serving/releases/tag/v0.5.0" class="markup--anchor markup--p-anchor" rel="noopener" target="_blank" >Serving</a >, <a href="https://github.com/knative/build/releases/tag/v0.5.0" data-href="https://github.com/knative/build/releases/tag/v0.5.0" class="markup--anchor markup--p-anchor" rel="noopener" target="_blank" >Build</a >, and <a href="https://github.com/knative/eventing/releases/tag/v0.5.0" data-href="https://github.com/knative/eventing/releases/tag/v0.5.0" class="markup--anchor markup--p-anchor" rel="noopener" target="_blank" >Eventing</a > repositories. </p>

<h3 name="813d" id="813d" class="graf graf--h3 graf-after--p"> Learn more </h3>

<ul class="postList">

<li name="c00d" id="c00d" class="graf graf--li graf-after--h3"> <a href="https://github.com/knative/docs#welcome-to-knative" data-href="https://github.com/knative/docs#welcome-to-knative" class="markup--anchor markup--li-anchor" rel="nofollow noopener noopener noopener noopener" target="_blank" >Welcome to Knative</a > </li>

<li name="d88a" id="d88a" class="graf graf--li graf-after--li"> <a href="https://knative.dev/docs/" target="_blank" >Getting started documentation</a > </li>

<li name="a9ca" id="a9ca" class="graf graf--li graf-after--li"> <a href="https://knative.dev/docs/samples/" target="_blank" >Samples and demos</a > </li>

<li name="fe75" id="fe75" class="graf graf--li graf-after--li"> <a href="https://knative.dev/contributing/" target="_blank" >Knative meetings and work groups</a > </li>

<li name="cf03" id="cf03" class="graf graf--li graf-after--li"> <a href="https://knative.dev/community/" target="_blank" >Questions and issues</a > </li>

<li name="19ce" id="19ce" class="graf graf--li graf-after--li"> Knative on Twitter (<a href="https://twitter.com/KnativeProject" data-href="https://twitter.com/KnativeProject" class="markup--anchor markup--li-anchor" rel="nofollow noopener noopener noopener noopener" target="_blank" >@KnativeProject</a >) </li>

<li name="5fcb" id="5fcb" class="graf graf--li graf-after--li"> Knative on <a href="https://stackoverflow.com/questions/tagged/knative" data-href="https://stackoverflow.com/questions/tagged/knative" class="markup--anchor markup--li-anchor" rel="nofollow noopener noopener noopener noopener" target="_blank" >StackOverflow</a > </li>

<li name="3813" id="3813" class="graf graf--li graf-after--li graf--trailing" > Knative <a href="https://slack.knative.dev/" data-href="https://slack.knative.dev/" class="markup--anchor markup--li-anchor" rel="nofollow noopener noopener noopener noopener" target="_blank" >Slack</a > </li>

</ul>

</div>

</div>

</section>

</section>

<footer>

<p> By <a href="https://medium.com/@mchmarny_google" class="p-author h-card" >Mark Chmarny</a > on <a href="https://medium.com/p/cfe646ca8e30" ><time class="dt-published" datetime="2019-04-11T14:30:40.853Z" >April 11, 2019</time ></a >. </p>

<p> <a href="https://medium.com/@mchmarny_google/announcing-knative-v0-5-release-cfe646ca8e30" class="p-canonical" >Canonical link</a > </p>

<p> Exported from <a href="https://medium.com">Medium</a> on June 7, 2019. </p>

</footer>

</article>
