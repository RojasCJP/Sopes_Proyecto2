FROM archlinux
COPY script.sh /script.sh
COPY Frontend /home/Frontend
RUN pacman -Syu --noconfirm
RUN pacman -S nodejs npm --noconfirm

ENTRYPOINT ["/script.sh"]
