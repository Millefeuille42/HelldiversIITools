# Helldivers II Discord Bot Deployment Manual

## Operation Briefing

Welcome, Helldiver! You're about to deploy the Helldivers II Discord Bot, your trusty companion in the ongoing war against the enemies of Super Earth. This bot is a crucial asset for staying updated on mission-critical information right within your Discord server. Here's your deployment manual:

The Helldivers II Discord Bot is accessible at [this url](https://discord.com/oauth2/authorize?client_id=1219964573231091713&permissions=277025705024&scope=bot) for immediate deployment.
Alternatively, it can be self-deployed using the provided resources.

## Mission Objectives

Your mission, should you choose to accept it, involves the following objectives:

1. **Install the Bot**: Deploy the bot in your Discord server to access its arsenal of commands.
2. **Configure Channels**: Set up the channels where the bot will relay vital intel and updates.
3. **Stay Informed**: Leverage the bot's commands to retrieve news, galaxy stats, major orders, and planetary details.
4. **Maintain Vigilance**: The bot will keep an eye on in-game events and provide alerts such as new major orders, feed messages, and planetary status changes.

## Reconnaissance

### Command Overview

Your bot is equipped with the following slash-commands:

- **news**: Acquire the latest news bulletin.
- **galaxy**: Gather intelligence on galaxy statistics.
- **order**: Receive the latest major order.
- **channel**: Designate a channel for the bot's transmissions.
- **planet**: Extract detailed data on specific planets.

### API Support

To fuel your operations, the bot taps into the official Helldivers II API.

## Deployment Tactics

### Components

Your arsenal includes the following components:

- **Bot**: The front-line interface with Discord servers.
- **API**: A covert transformer/caching hub.
- **Updater**: The silent sentinel, notifying events to the bot.
- **CLI**: Command-line interface for manual interactions.

### Logistics

Logistical details to keep in mind:

- **Cache**: Utilize Redis caching for swift data retrieval.
- **Assembly**: Locate and build binaries in the `cmd/<binary>` directory with the `go build cmd/<binary>` stratagem.
- **Fortification**: Strengthen your defenses with Docker. Each component has a dedicated Dockerfile (`Dockerfile_<binary>`).
- **Formation**: Assemble your forces with the provided Docker Compose file for production.
- **Configuration**: Required parameters and variables are detailed in `sample.env`

## Rules of Engagement

### Legal Clearance

This mission operates under the MIT License. 
Remember, while our cause aligns with the ideals of Super Earth, we are independent operators. 
In case of inquiries or disputes, refer to the designated contact.

### Usage Policies

In the event you would use the pre-deployed version, data collected, being:
- commands used (excluding origin of the command)
- server ID
- server name
- channel ID

Is solely for debugging and troubleshooting purposes. 
No individual, including the operator, will access to logs and data for any other purposes.

### Contact

For further briefing or assistance, contact your mission coordinator [Millefeuille](mailto:millefeuille42@proton.me).

## For Super Earth!

Your commitment to the cause is commendable. With the Helldivers II Discord Bot at your side, victory is within reach. 
Deploy the bot, rally your comrades, and together, we shall triumph against the enemies of Super Earth!
