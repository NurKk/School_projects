#!/bin/bash
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"Nurzhol\"}}){id}}"}' | grep -o '[0-9]*'


