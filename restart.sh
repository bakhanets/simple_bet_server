docker stop bet-service
docker compose up -d --build

https://download.maxmind.com/app/geoip_download_by_token?edition_id=GeoLite2-Country&date=20221011&suffix=tar.gz&token=v2.local.uWyRzIZOMVB41Zt-WpY5N9bC27tcFaSJx-SRm74UKfpsNCskVssWD0tkUdJ_nqqkpQ7mFpDEXw2EKfe6uAeacmZQqAqfc8qELbEGPe1qBQUemEG-JJxtZrVTzx56txH2aXBhPQsgWKsijfNZgK86fHTWivO56PLnrOSTgF5zS1Z8tYlv11ysFAYesGY2jS-rF1tkcQ
curl "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=HhJ3cMvkJxbo3NGN&suffix=tar.gz" -o GeoLite2-Country.tar.gz \
  && tar -xzvf GeoLite2-Country.tar.gz \
  && mkdir /geoip
  && mv GeoLite2-Country_*/GeoLite2-Country.mmdb /geoip/GeoLite2-Country.mmdb