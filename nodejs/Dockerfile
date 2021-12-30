FROM archlinux
COPY script.sh /script.sh
COPY Frontend /home/Frontend
RUN pacman -Syu --noconfirm
RUN pacman -S nodejs npm --noconfirm
RUN pacman -S nano --noconfirm


ENTRYPOINT ["/script.sh"]