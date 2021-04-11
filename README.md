# Nat20
## D&D For Busy People

### Purpose
Nat20 is a discord bot developed with the sole purpose of running asynchronous D&D campaigns.

### Team Member and Roles
Evan Mills: Discord Bot Interface
Max Farver: Backend and Web UI

### Approach
The approach taken was to have a bot control a server and manage text channels and roles to simulate traveling through the
world of a D&D campaign. How it works is after getting added to the server the player who is going to be the DM starts the 
campaign via the bot command. This will create a session and allow for the registration of players. Once the session is created
the DM will be given access to the admin page where they can add various locations and events to build out a robust world. 

Once the players register and create their character they can choose a stating location. This is done through the by declaring
the location in which they want to start in. From there they are open to explore the world that is built around them. 

To simulate moving around the world the bot is set up to create and grant access to role specific text channels. By creating 
these channels to be role specific only people in that location can view what is happening there. 

When an interaction occurs a separate text channel will be set up for the DM to provide the necessary responses to players
actions.

### Post HackKU
After finishing HackKU we plan to develop the project further. We plan on adding in a way for players to manage their inventory,
along with streamlining the more basic events and interactions.
