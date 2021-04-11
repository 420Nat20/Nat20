import os
from random import randint

import discord
from discord.ext import commands
from dotenv import load_dotenv
import requests

PREFIX = '!'

f = open('commands.md')
command_message = f.read()

load_dotenv()
TOKEN = os.getenv('TOKEN')

intents = discord.Intents.default()
intents.members = True
client = commands.Bot(command_prefix=PREFIX, intents=intents)

base_endpoint_url = 'https://api.nat20.tech'


@client.event
async def on_ready():
    await client.change_presence(activity=discord.Game(name='!commands to list possible commands'))
    print('Online')


@client.command()
async def temp(ctx):
    nat20 = discord.utils.get(ctx.message.guild.roles, name='Nat20')
    await nat20.edit(position=2)


@client.command()
async def start_campaign(ctx):
    author = ctx.message.author

    response = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/')
    print(response.text)
    if response.text.strip() != 'record not found':
        await ctx.message.channel.send(f'<@{author.id}>, a game is already running wait until that one is done.')
        return

    dm = discord.utils.get(ctx.message.guild.roles, name='DM')
    if dm is None:
        dm = await ctx.message.guild.create_role(name='DM', permissions=discord.Permissions.all())
        await author.add_roles(dm)
        await create_channel(ctx, 'DM')

    body = {'ServerID': ctx.message.guild.id, 'DM': author.id}
    requests.post(f'{base_endpoint_url}/games/', json=body)


@client.command()
async def end_campaign(ctx):
    author = ctx.message.author
    dm = discord.utils.get(ctx.message.guild.roles, name='DM')

    if dm not in author.roles:
        await ctx.message.channel.send(f'{author.mention}, You are not the DM and cannot end the campaign.')
        return

    response = requests.delete(f'{base_endpoint_url}/games/{ctx.message.guild.id}/')
    print(response.text)
    if response.text.strip() == 'record not found':
        await ctx.message.channel.send(f'<@{author.id}>, there is not a game running.')
        return

    roles = ctx.message.guild.roles
    for role in roles:
        if role.name != 'Nat20' and role.name != '@everyone':
            await role.delete()

    channels = ctx.message.guild.text_channels
    for channel in channels:
        if channel.name != 'general':
            await channel.delete()


@client.command()
async def character(ctx):
    author = ctx.message.author
    stats = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/users/{author.id}')

    if stats.text == 'record not found':
        await ctx.message.channel.send(f'<@{author.id}, You need to register your character. You can do this by using '
                                       f'`!register`.')
        return

    wanted = ['Name', 'Class', 'Background', 'Race', 'Alignment', 'Strength', 'Dexterity', 'Constitution',
              'Intelligence', 'Wisdom', 'Charisma', 'Ideal', 'TraitOne', 'TraitTwo', 'Bond', 'Flaw']

    keys = stats.keys()
    stat_str = 'Here is you current character sheet:\n'
    for key in keys:
        if key in wanted:
            if key == 'TraitOne':
                stat_str += f'> Trait One: {stats[key]}\n'
            elif key == 'TraitTwo':
                stat_str += f'> Trait Two: {stats[key]}\n'
            else:
                stat_str += f'> {key}: {stats[key]}\n'

    await author.send(stat_str)


@client.command()
async def roll(ctx):
    """
    Command used for rolling dice
    :param ctx: context for the command
    :return: None
    """
    valid_dice = ['4', '6', '8', '10', '12', '20']
    author = ctx.message.author
    message = ctx.message.content
    message = message.replace('!roll ', '')

    split = message.split('x')

    if len(split) == 1:
        split[0] = split[0].replace('d', '')
        if split[0] not in valid_dice:
            await ctx.message.channel.send(f'<@{author.id}> Your input was invalid input for the roll please try again')
            return
    elif len(split) == 2:
        split[1] = split[1].replace('d', '')
        if split[1] not in valid_dice:
            await ctx.message.channel.send(f'<@{author.id}> Your input was invalid input for the roll please try again')
            return
    else:
        await ctx.message.channel.send(f'<@{author.id}> Your input was invalid input for the roll please try again')
        return

    for i in range(len(split)):
        split[i] = int(split[i])

    sum = 0
    if len(split) == 1:
        sum = randint(1, split[0] + 1)
    else:
        for i in range(split[0]):
            sum += randint(1, split[1] + 1)

    await ctx.message.channel.send(f'<@{author.id}>, The sum of your {message} is {sum}')


@client.command()
async def commands(ctx):
    await ctx.message.channel.send(command_message)


@client.command()
async def register(ctx):
    """
    Command used to register a character
    :param ctx: context for the command
    :return: None
    """
    await ctx.message.author.send(f'Follow this link to register your character.')


@client.command()
async def locations(ctx):
    """
    Command to list out the locations in the world
    :param ctx: context for the command
    :return: None
    """
    locations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild}/locations/')
    locations = locations.json()
    channel = ctx.message.channel

    location_string = f'<@{ctx.message.author.id}> Here are the available locations:\n'

    for location in locations:
        name = location['Name']
        location_string += f'> {name} \n'

    await channel.send(location_string)


@client.command()
async def sublocations(ctx):
    author = ctx.message.author
    author_roles = author.roles
    location = None

    for temp in author_roles:
        if '_' not in temp.name and temp.name != '@everyone' and temp.name != 'DM':
            location = temp
            break

    location_id = None
    locations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/').json()
    for loc in locations:
        if loc['Name'] == location.name:
            location_id = loc['ID']
            break

    sublocations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/{location_id}/sublocations/').json()
    channel = ctx.message.channel

    location_string = f'<@{ctx.message.author.id}> Here are the available sublocations at {location.name}:\n'

    for location in sublocations:
        name = location['Name']
        location_string += f'> {name} \n'

    await channel.send(location_string)


@client.command()
async def end_encounter(ctx):
    author = ctx.message.author
    dm = discord.utils.get(author.roles, name='DM')

    if not dm:
        await ctx.message.channel.send(f'<@{author.id}>, You do not have permission to end encounters')
        return

    members = ctx.message.channel.members

    for member in members:
        roles = member.roles
        for role in roles:
            if 'encounter' in role.name:
                await member.remove_roles(role)

    await ctx.message.channel.delete()


@client.command()
async def travel(ctx):
    author = ctx.message.author
    message = ctx.message.content

    location = message.replace('!travel ', '')

    locations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/').json()

    valid = False
    location_data = None
    for loc in locations:
        if loc['Name'] == location:
            valid = True
            location_data = loc
            break

    if not valid:
        await ctx.message.channel.send(f'<@{author.id}>, That is not a valid location. Use `!locations` '
                                       f'to list out possible values.')
        return

    author_roles = author.roles
    for role in author_roles:
        if role.name != '@everyone' and role.name != 'DM':
            await author.remove_roles(role)
            if len(role.members) == 0:
                await role.delete()
                if '_' in role.name:
                    if '_' in ctx.message.channel.name:
                        await ctx.message.channel.delete()
                    else:
                        channel = discord.utils.get(ctx.message.guild.text_channels, name=role.name.replace(' ', '-'))
                        if channel is not None:
                            await channel.delete()

    location_role = discord.utils.get(ctx.message.guild.roles, name=location)

    if location_role is None:
        await create_role(ctx, location)
        await create_role(ctx, f'{location}_general')
        await create_channel(ctx, f'{location}_general')
    else:
        general_role = discord.utils.get(ctx.message.guild.roles, name=f'{location}_general')

        if general_role is None:
            create_role(ctx, f'{location}_general')
            await create_channel(ctx, f'{location}_general')
        else:
            await author.add_roles(general_role)

        await author.add_roles(location_role)

    print(location_data)
    if not location_data['Visited'] and location_data['EventDescription']:
        location_data['Visited'] = True
        print(location_data)
        loc_id = location_data['ID']
        requests.put(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/{loc_id}',
                     json=location_data)
        await create_role(ctx, f'{location}_encounter')
        encounter_channel = await create_channel(ctx, f'{location}_encounter')
        dm = discord.utils.get(ctx.message.guild.roles, name='DM')
        await encounter_channel.send(f'<@{dm}>, There is an encounter starting.')
        await encounter_channel.send(location_data['EventDescription'])


@client.command()
async def enter(ctx):
    """
    Enter a sublocation
    :param ctx: context of the message
    :return: None
    """
    author = ctx.message.author
    message = ctx.message.content

    sublocation = message.replace('!enter ', '')

    author_roles = author.roles
    location = None

    for temp in author_roles:
        if '_' not in temp.name and temp.name != '@everyone' and temp.name != 'DM':
            location = temp
            break

    location_id = None
    locations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/').json()
    for loc in locations:
        if loc['Name'] == location.name:
            location_id = loc['ID']
            break

    valid = False
    sublocations = requests.get(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/{location_id}/sublocations/').json()
    sublocation_data = None
    for sub in sublocations:
        if sub['Name'] == sublocation:
            valid = True
            sublocation_data = sub

    if not valid:
        await ctx.message.channel.send(f'<@{author.id}>, That is not a valid sublocation. Use `!sublocations` '
                                       f'to list out possible values.')
        return

    role = discord.utils.get(ctx.message.guild.roles, name=f'{location.name}_{sublocation}')

    for temp in author_roles:
        if '_' in temp.name:
            await author.remove_roles(temp)
            if len(temp.members) == 0:
                if '_' == ctx.message.channel.name:
                    await ctx.message.channel.delete()
                else:
                    channel = discord.utils.get(ctx.message.guild.text_channels, name=temp.name.replace(' ', '-'))
                    if channel is not None:
                        await channel.delete()
                await temp.delete()

    if role is None:
        await create_role(ctx, f'{location.name}_{sublocation}')
        await create_channel(ctx, f'{location.name}_{sublocation}')
    else:
        await author.add_roles(role)

    if not sublocation_data['Visited'] and sublocation_data['EventDescription']:
        sublocation_data['Visited'] = True
        sub_id = sublocation_data['ID']
        requests.put(f'{base_endpoint_url}/games/{ctx.message.guild.id}/locations/{location_id}/sublocations/{sub_id}',
                     json=sublocation_data)
        await create_role(ctx, f'{location.name}_{sublocation}_encounter')
        encounter_channel = await create_channel(ctx, f'{location.name}_{sublocation}_encounter')
        dm = discord.utils.get(ctx.message.guild.roles, name='DM')
        await encounter_channel.send(f'<{dm.mention}>, There is an encounter starting.')
        await encounter_channel.send(sublocation_data['EventDescription'])


@client.command()
async def leave(ctx):
    author = ctx.message.author

    roles = ctx.message.guild.roles
    author_roles = author.roles
    location = None

    for temp in roles:
        if '_' not in temp.name and temp.name != '@everyone' and temp.name != 'DM':
            location = temp
            break

    role = discord.utils.get(ctx.message.guild.roles, name=f'{location.name}_general')

    for temp in author_roles:
        if '_' in temp.name:
            await author.remove_roles(temp)
            if len(temp.members) == 0:
                if '_' == ctx.message.channel.name:
                    await ctx.message.channel.delete()
                else:
                    channel = discord.utils.get(ctx.message.guild.text_channels, name=temp.name.replace(' ', '-'))
                    if channel is not None:
                        await channel.delete()
                await temp.delete()

    if role is None:
        await create_role(ctx, f'{location.name}_general')
        await create_channel(ctx, f'{location.name}_general')
    else:
        await author.add_roles(role)


async def create_channel(ctx, channel_name):
    """
    Function to create and manage a new channel
    :param ctx: context of channel create
    :param channel_name: name for the channel to be created
    :return: None
    """
    role = discord.utils.get(ctx.message.guild.roles, name=channel_name)

    if role is None:
        print('Role not found')
        return

    overwrites = {
        ctx.message.guild.default_role: discord.PermissionOverwrite(read_messages=False),
        ctx.message.guild.me: discord.PermissionOverwrite(read_messages=True),
        role: discord.PermissionOverwrite(read_messages=True, send_messages=True)
    }
    return await ctx.message.guild.create_text_channel(channel_name, overwrites=overwrites)


async def create_role(ctx, role_name):
    role = await ctx.message.guild.create_role(name=role_name)
    await ctx.message.author.add_roles(role)


client.run(TOKEN)
