import asyncio
import os
from random import randint

import discord
from discord.ext import commands
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv('TOKEN')
client = commands.Bot(command_prefix='!')


@client.event
async def on_ready():
    print('Online')


@client.command()
async def roll(ctx):
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

client.run(TOKEN)