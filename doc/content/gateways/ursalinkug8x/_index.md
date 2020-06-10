---
title: "Ursalink UG8X IoT LoRaWAN Gateway"
description: ""
---

The **Ursalink UG8X IoT LoRaWAN Gateway** is an 8 channel (16 channel optional) configurable, scalable gateway for industrial IoT applications.

This page contains information about connecting the Ursalink UG8X IoT LoRaWAN Gateway to {{% tts %}}

<!--more-->

The technical specifications can be found in [Ursalink's official documentation](https://www.ursalink.com/en/ad-lorawan-gateway/). The Ursalink UG8X IoT LoRaWAN Gateway supports two ways to connect with {{% tts %}}, using either the Semtech Packet Router or LoRa Basic Station.

{{< figure src="ursalink.jpg" alt="Ursalink" >}}

## Requirements

1. User account on {{% tts %}} with rights to create gateways.
2. Ursalink UG8X LoRaWAN Gateway connected to the internet via ethernet or cellular.
3. CA certificate for the Basic Station (if using Basic Station).

## Registration

Create a gateway by following the instructions for the [Console]({{< ref "/getting-started/console#create-gateway" >}}) or the [CLI]({{< ref "/getting-started/cli#create-gateway" >}}).

The **EUI** of the gateway can be found on the configuration web page of the gateway.

{{< figure src="eui.png" alt="Gateway EUI" >}}

## Configuration via Browser

Find the IP address of the gateway. The default IP for the Ursalink UG8X LoRaWAN Gateway is 192.168.23.150.

Connect your machine to the same local network as that of the gateway, and enter the IP address in your web browser. The default username is **admin** and the default password is **password**. See [Ursalink's official documentation](https://www.ursalink.com/en/ad-lorawan-gateway/) for more information.

{{< figure src="login.png" alt="Login" >}}

### Disable Default Server

In the left menu, choose **Packet Forwarder**. Select the **General** tab.

{{< figure src="eui.png" alt="Packet Forwarder" >}}

Click the pencil icon next to the default server, and uncheck the **Enabled** button to disable the default server.

Click **Save** to continue.

{{< figure src="disable.png" alt="Disable default server" >}}

## Configuring the Semtech Packet Router

In the **Packet Forwarder** menu, click the **Plus** button to create a new server.

{{< figure src="plus.png" alt="Create new server" >}}

In the server configuration options, check the **Enabled** box.

Choose **Semtech** as the **Type**.

For the **Server Address** choose **custom**, and enter the same as what you use instead of `thethings.example.com` in the [Getting Started guide]({{< ref "/getting-started" >}}).

Choose the appropriate **Port Up** and **Port Down** values. These are both **1700** by default in {{% tts %}}.

Click **Save** to continue.

{{< figure src="semtech.png" alt="Semtech Configuration" >}}

If your configuration was successful, your gateway will connect to {{% tts %}} after a couple of seconds.

## Configuring Basic Station

In the **Packet Forwarder** menu, click the **Plus** button to create a new server.

{{< figure src="plus.png" alt="Create new server" >}}

In the server configuration options, check the **Enabled** box.

Choose **Basic Station** as the **Type**.

Enter the **URI**, **CA File**, **Client Certificate**, and **Client Key File** from your CUPS and LNS configuration information.

Click **Save** to continue.

{{< figure src="basic-station.png" alt="Basic Station Configuration" >}}

If your configuration was successful, your gateway will connect to {{% tts %}} after a couple of seconds.
