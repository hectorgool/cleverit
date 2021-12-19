# start with base image
FROM mysql:8.0.27

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
#ADD ./schema/data.sql /docker-entrypoint-initdb.d/

#RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

#CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]
