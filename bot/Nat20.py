import asyncio
import os
from random import randint

import discord
from discord.ext import commands
from dotenv import load_dotenv


PREFIX='!'

f = open('commands.md')
command_message = f.read()

load_dotenv()
TOKEN = os.getenv('TOKEN')
client = commands.Bot(command_prefix=PREFIX)


@client.event
async def on_ready():
    await client.change_presence(activity=discord.Game(name='!commands to list possible commands'))
    print('Online')


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
    pass


@client.command()
async def locations(ctx):
    """
    Command to list out the locations in the world
    :param ctx: context for the command
    :return: None
    """
    channel = ctx.message.channel
    await channel.send('Locations')


@client.command()
async def travel(ctx):
    author = ctx.message.author
    print(author)
    message = ctx.message.content

    location = message.replace('!travel ', '')

    roles = ctx.message.guild.roles

    location_role = None
    for temp in roles:
        if temp.name == location:
            location_role = temp
            break

    if location_role is None:
        await create_role(ctx, location)
        await create_role(ctx, f'{location}-general')
        await create_channel(ctx, f'{location}-general')
    else:
        general_role = None
        for temp in roles:
            if temp.name == f'{location}-general':
                general_role = temp
                break

        if general_role is None:
            create_role(ctx, f'{location}-general')
            await create_channel(ctx, f'{location}-general')
        else:
            await ctx.message.author.add_roles(general_role)

        await ctx.message.author.add_roles(location_role)


@client.command()
async def enter(ctx):
    pass


async def create_channel(ctx, channel_name):
    """
    Function to create and manage a new channel
    :param ctx: context of channel create
    :param channel_name: name for the channel to be created
    :return: None
    """
    roles = ctx.message.guild.roles
    role = None
    for temp in roles:
        if temp.name == channel_name:
            role = temp
            break

    if role is None:
        print('Role not found')
        return

    print(role)
    overwrites = {
        ctx.message.guild.default_role: discord.PermissionOverwrite(read_messages=False),
        ctx.message.guild.me: discord.PermissionOverwrite(read_messages=True),
        role: discord.PermissionOverwrite(read_messages=True, send_messages=True)
    }
    await ctx.message.guild.create_text_channel(channel_name, overwrites=overwrites)


async def create_role(ctx, role_name):
    role = await ctx.message.guild.create_role(name=role_name)
    await ctx.message.author.add_roles(role)

client.run(TOKEN)