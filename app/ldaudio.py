import os
import subprocess
import runpy
import time
import lunlib

lunlib.loading()

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

# exec = 'ffmpeg'
# inputfile = 'c:\'
# inputfile1 = 'c:\'
# inputfile2 = 'c:\'
# optlist = f'{exec} {inputfile} {inputfile} {inputfile} {output}'
# ffmpeg = exec + optlist
# print(ffmpeg)
# opts = ''
# subprocess.run(f'ffmpeg {opts}')