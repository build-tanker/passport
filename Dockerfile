FROM scratch
LABEL maintainer="sudhanshu@go-jek.com"
ADD bin/passport_linux passport
ENV PORT 3000
EXPOSE 3000
ENTRYPOINT ["/passport", "start"]