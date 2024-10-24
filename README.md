# Go-Go-Go
Trabajo Practico para la materia "Teoría de lenguajes" de la Facultad de Ingeniería de la Universidad de Buenos Aires.

Nuestro proyecto va a ser una app de consulta y estadisticas del mercado de bitCoins, desde la cual se podrán ver distintos gráficos, datos y cotizaciones segun la moneda seleccionada, conversiones entre las distintas monedas y tambien alertas de precios segun superen algun umbral establecido por el usuario.
Para esto vamos a desarrollar una RestAPI con distintos endpoints que a su vez funcionaran como Proxy de la API CoinGecko la cual nos proveera de toda la información que necesitemos.
Para la UI vamos a usar el paquete de GO html/template con el cual podemos generar plantillas de html dinámicas para nuestro servidor.

Las principales Features que va a tener van a ser:
Estadísticas agregadas: Promedio de precio, precio máximo/mínimo en un periodo. Hacer Graficos para el precio a lo largo del tiempo.
Conversión de monedas: Entre Bitcoin, Ethereum, y otras criptomonedas o monedas fiduciarias (USD, EUR, etc.). (Pensar en agregar mas cosas)
Alertas de precios: Posibilidad de configurar un endpoint para que la API envíe alertas cuando el precio de una criptomoneda suba o baje de cierto umbral.
Información del mercado: Volumen de transacciones, capitalización de mercado de las monedas, entre otros.
