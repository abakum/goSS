cd /d "%~dp0"
if exist %~n0.log goto :EOF
echo %date% %time% %USERNAME%>%~n0.log
:call :crop doc\01.png 01.jpg "-vf crop=881:413:503:296"
:call :crop doc\04.png 04.jpg "-vf crop=994:443:254:171"
:call :crop doc\05.png 05.jpg "-vf crop=1488:656:272:134"
:call :pdf2jpg 08 "-vf crop=1754:1111:0:64"
set ex=mp4
set sec=40
:call :fd %Date:~3,2%%Date:~0,2%
xcopy /DYLR *.jpg %~n0.%ex%*|findstr /BC:"0 ">nul||call :jpg2mov
del %~n0.log
goto :EOF

:crop 
 if not exist %1 goto :EOF
 xcopy /DYLR %1 %2*|findstr /BC:"0 ">nul&&goto :EOF
 bin\ffmpeg.exe -i %1 %~3 -q:v 1 -y %2
goto :EOF

:pdf2jpg
 xcopy /DYLR doc\%1.pdf %1-000001.jpg*|findstr /BC:"0 ">nul&&goto :EOF
 del doc\%1-*.png
 del %1-*.jpg
 bin\pdftopng.exe doc\%1.pdf doc\%1
 call :crop doc\%1-000001.png %1-000001.jpg %2
 :call :crop doc\%1-000002.png %1-000002.jpg %2
 del doc\%1-*.png
goto :EOF

:jpg2mov
 type *.jpg|bin\ffmpeg.exe -framerate 1/%sec% -f image2pipe -i - -vf "scale=1888:1888*9/16:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2" -c:v libx264 -tune stillimage -r 25 -y %~n0.%ex%
 :type *.jpg|bin\ffmpeg.exe -framerate 1/%sec% -f image2pipe -i - -vf "scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2" -r 1/%sec% -y %~n0.apng
 :ffmpeg -i %~n0.%ex% -filter_complex "fps=1/%sec%,split[v1][v2]; [v1]palettegen=stats_mode=full [palette]; [v2][palette]paletteuse=dither=sierra2_4a" -r 1/%sec% -y %~n0.gif
 :ffmpeg -i %~n0.%ex% -filter_complex "fps=1/%sec% -r 1/%sec% -y %~n0.apng
 :type *.jpg|bin\ffmpeg.exe -framerate 1 -f image2pipe -i - -filter_complex "fps=1,scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,split[v1][v2]; [v1]palettegen=stats_mode=full [palette]; [v2][palette]paletteuse=dither=sierra2_4a" -r 1/%sec% -y abaku.gif
 echo %date% %time% %USERNAME%>>doc\%~n0.log
 start %~n0.%ex%
goto :EOF

:fd
 set d=%1
 if not exist fd\99%d%-1.jpg set d=0631
 if exist 99%d%-1.jpg goto :EOF
 del 99*.jpg
 xcopy fd\99%d%-?.jpg
 copy /b 99%d%-1.jpg +,,
goto :EOF

