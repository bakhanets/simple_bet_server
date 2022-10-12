FROM golang:1.18 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v

RUN CGO_ENABLED=0 go build -o /go/bin/app
RUN curl "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=HhJ3cMvkJxbo3NGN&suffix=tar.gz" -o GeoLite2-Country.tar.gz && tar -xzvf GeoLite2-Country.tar.gz

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
COPY --from=build /go/src/app/GeoLite2-Country_*/GeoLite2-Country.mmdb /geoip/GeoLite2-Country.mmdb
CMD ["/app"]