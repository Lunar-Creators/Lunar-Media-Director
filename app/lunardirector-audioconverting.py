import os
import subprocess
import runpy
import time

def loading():
    print('\033[3m\033[33m{}\033[0m'.format('Loading...'))
loading()

def inputfile():
    from tkinter import filedialog as seele

    filetypes = (
            ('Video or Audio files',['*.avi','*.mp4','*.mkv','*.mov','*.vob','*.flv','*.wmv','*.3gp','*.pcm','*.wav','*.flac','*.mp3','*.aac','*.ogg','*.alac','*.mka','*.m4a','*.oga','*.opus','*.wma','*.webm']),
            ('Audio',['*.pcm','*.wav','*.flac','*.mp3','*.aac','*.ogg','*.alac','*.mka','*.m4a','*.oga','*.opus','*.wma','*.webm']),
            ('Video',['*.avi','*.mp4','*.mkv','*.mov','*.vob','*.flv','*.webm','*.wmv','*.3gp']),
            ('All files', '*.*')
        )
    fileselect = seele.askopenfilename(
            title='Select file',
            initialdir='/',
            filetypes=filetypes)
    opts = '-i' + f' "{fileselect}"'

# Continue working with aPresets
# opts = f'{opts}' + f' -c:a {variable}'

def wronginput():
    print('\033[1m\033[91m{}\033[0m'.format('INCORRECT INPUT'))
    time.sleep(1)
    os.system('cls' if os.name == 'nt' else 'clear')

def leavemodule():
    print('\033[3m\033[33m{}\033[0m'.format('Leaving...'))

# exec = 'ffmpeg'
# inputfile = 'c:\'
# inputfile1 = 'c:\'
# inputfile2 = 'c:\'
# optlist = f'{exec} {inputfile} {inputfile} {inputfile} {output}'
# ffmpeg = exec + optlist
# print(ffmpeg)
# opts = ''
# subprocess.run(f'ffmpeg {opts}')