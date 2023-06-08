import os
import subprocess
import time

def loading():
    print('\033[3m\033[33m{}\033[0m'.format('Loading...'))

loading()

def wronginput():
    print('\033[1m\033[91m{}\033[0m'.format('INCORRECT INPUT'))
    time.sleep(1)
    os.system('cls' if os.name == 'nt' else 'clear')

def leavemodule():
    print('\033[3m\033[33m{}\033[0m'.format('Leaving...'))
# Universal part

def helpmenu():
    print('\033[1m{}\033[0m'.format('\nHow can ffmpeg help you?'),'\n--------------------------\n  1 - Standart list of help\n  2 - Very long List of help\n  3 - EXTREAMLY LONG LIST OF HELP\n  4 - Flag help for available ffmpeg encoders\n  5 - List of available ffmpeg encoders\n  Q - Quit Help\n--------------------------\n','\033[35m{}\033[0m'.format('NOTE! In commandline mode, you can enter the help command with your arguments'))
    choice = input("Type your choice: ")
    if choice == '1':
        subprocess.run('ffmpeg -h', check=False)
        helpmenu()
    if choice == '2':
        subprocess.run('ffmpeg -h long', check=False)
        helpmenu()
    if choice == '3':
        subprocess.run('ffmpeg -h full', check=False)
        helpmenu()
    if choice == '4':
        subprocess.run('ffmpeg -encoders', check=False)
        print('\033[1m{}\033[0m'.format('\n\nEnter the one you are interested in from the list'),'\033[33m{}\033[0m'.format('\n(example: libvpx-vp9)'),)
        encoder = input("You need help with: ")
        subprocess.run(f'ffmpeg -h encoder={encoder}', check=False)
        del encoder
        helpmenu()
    if choice == '5':
        subprocess.run('ffmpeg -encoders', check=False)
        helpmenu()
    if choice == 'Q':
        leavemodule()
    else:
        wronginput()
        helpmenu()
os.system('cls' if os.name == 'nt' else 'clear')
helpmenu()    

