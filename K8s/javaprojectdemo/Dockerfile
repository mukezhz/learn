# Set a variable that can be used in all stages.
ARG BUILD_HOME=/gradle

# Gradle image for the build stage.
FROM gradle:jdk17-alpine as build-image

# Set the working directory.
ARG BUILD_HOME
ENV APP_HOME=$BUILD_HOME
WORKDIR $APP_HOME

# Copy the Gradle config, source code, and static analysis config
# into the build container.
COPY --chown=gradle:gradle app/build.gradle settings.gradle $APP_HOME/
COPY --chown=gradle:gradle app/src $APP_HOME/src
# COPY --chown=gradle:gradle config $APP_HOME/config

# Build the application.
RUN gradle --no-daemon build
# Java image for the application to run in.
FROM openjdk:17-alpine

# Copy the jar file in and name it app.jar.
ARG BUILD_HOME
ENV APP_HOME=$BUILD_HOME
WORKDIR $APP_HOME
COPY --from=build-image $APP_HOME/build/libs/*.jar app.jar

# The command to run when the container starts.
# ENTRYPOINT java -jar app.jar
CMD [ "java", "-jar", "app.jar" ]