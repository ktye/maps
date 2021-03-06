package tile

// Cities is a collection of cities and their coordinates,
// which have a population of more than 1E6 people.
// Reference:
// https://www.maxmind.com/en/free-world-cities-database
// License: http://download.maxmind.com/download/geoip/database/LICENSE_WC.txt
// "This product includes data created by MaxMind, available from http://www.maxmind.com/"
var Cities = map[string]LatLon{
	"abidjan":             {5.309657, -4.012656},
	"accra":               {5.55, -0.2166667},
	"adana":               {37.001667, 35.328889},
	"addis abeba":         {9.024325, 38.749226},
	"adelaide":            {-34.928661, 138.598633},
	"agra":                {27.183333, 78.016667},
	"ahmadabad":           {23.033333, 72.616667},
	"aleppo":              {36.2027778, 37.1586111},
	"alexandria":          {31.1980556, 29.9191667},
	"algiers":             {36.7630556, 3.0505556},
	"allahabad":           {25.45, 81.85},
	"almaty":              {43.25, 76.95},
	"amritsar":            {31.633056, 74.865556},
	"ankara":              {39.911652, 32.840305},
	"anshan":              {41.123611, 122.99},
	"antananarivo":        {-18.9166667, 47.5166667},
	"aurangabad":          {19.883333, 75.333333},
	"baghdad":             {33.3386111, 44.3938889},
	"baku":                {40.395278, 49.882222},
	"bamako":              {12.65, -8.0},
	"bandung":             {-6.903889, 107.618611},
	"bangalore":           {12.983333, 77.583333},
	"bangkok":             {13.753979, 100.501444},
	"barcelona":           {41.398371, 2.1741},
	"barranquilla":        {10.963889, -74.796389},
	"bayrut":              {33.8719444, 35.5097222},
	"beirut":              {33.8719444, 35.5097222},
	"bekasi":              {-6.2349, 106.9896},
	"belem":               {-1.45, -48.483333},
	"belgrade":            {44.818611, 20.468056},
	"belo horizonte":      {-19.916667, -43.933333},
	"benin":               {6.335045, 5.627492},
	"berlin":              {52.516667, 13.4},
	"bhopal":              {23.266667, 77.4},
	"bogota":              {4.649178, -74.062827},
	"bombay":              {18.975, 72.825833},
	"brasilia":            {-15.783333, -47.916667},
	"brazzaville":         {-4.2591667, 15.2847222},
	"brisbane":            {-27.47101, 153.024292},
	"brussels":            {50.833333, 4.333333},
	"bucharest":           {44.433333, 26.1},
	"budapest":            {47.5, 19.083333},
	"bursa":               {40.191667, 29.061111},
	"cairo":               {30.05, 31.25},
	"calcutta":            {22.569722, 88.369722},
	"cali":                {3.437222, -76.5225},
	"campinas":            {-22.9, -47.083333},
	"cape town":           {-33.925839, 18.423218},
	"caracas":             {10.5, -66.9166667},
	"casablanca":          {33.592779, -7.619157},
	"changchun":           {43.88, 125.322778},
	"chelyabinsk":         {55.154444, 61.429722},
	"chengdu":             {30.666667, 104.066667},
	"chicago":             {41.8500000, -87.6500000},
	"chongqing":           {29.562778, 106.552778},
	"conakry":             {9.5091667, -13.7122222},
	"copenhagen":          {55.666667, 12.583333},
	"cordoba":             {-31.413496, -64.181052},
	"curitiba":            {-25.416667, -49.25},
	"dakar":               {14.6708333, -17.4380556},
	"dalian":              {38.912222, 121.602222},
	"dallas":              {32.7833333, -96.8000000},
	"damascus":            {33.5, 36.3},
	"dar es salaam":       {-6.8, 39.2833333},
	"datong":              {40.093611, 113.291389},
	"davao":               {7.073056, 125.612778},
	"delhi":               {28.666667, 77.216667},
	"depok":               {-6.343333, 106.498889},
	"dhaka":               {23.7230556, 90.4086111},
	"douala":              {4.0502778, 9.7},
	"dubai":               {25.258172, 55.304717},
	"dublin":              {53.3330556, -6.2488889},
	"durban":              {-29.857896, 31.029198},
	"ecatepec":            {19.601111, -99.0525},
	"esfahan":             {32.657218, 51.677608},
	"faisalabad":          {31.416667, 73.083333},
	"faridabad":           {28.433333, 77.316667},
	"fez":                 {34.037151, -4.999797},
	"fortaleza":           {-3.316667, -41.416667},
	"freetown":            {8.49, -13.2341667},
	"fushun":              {41.7, 123.883333},
	"fuzhou":              {26.061389, 119.306111},
	"gaziantep":           {37.059444, 37.3825},
	"ghaziabad":           {28.666667, 77.433333},
	"gizeh":               {30.0086111, 31.2122222},
	"goiania":             {-16.666667, -49.266667},
	"guadalajara":         {20.666667, -103.333333},
	"guangzhou":           {23.116667, 113.25},
	"guarulhos":           {-23.45068, -46.526175},
	"guayaquil":           {-2.1666667, -79.9},
	"guiyang":             {26.583333, 106.716667},
	"gujranwala":          {32.161667, 74.188309},
	"ha noi":              {21.033333, 105.85},
	"hamburg":             {53.55, 10},
	"handan":              {36.566667, 114.533333},
	"hangzhou":            {30.29365, 120.161419},
	"hanoi":               {21.033333, 105.85},
	"haora":               {22.589167, 88.310278},
	"harare":              {-17.8177778, 31.0447222},
	"harbin":              {45.75, 126.65},
	"hefei":               {31.863889, 117.280833},
	"hiroshima":           {34.434844, 132.740622},
	"ho chi minh city":    {10.75, 106.666667},
	"houston":             {29.7630556, -95.3630556},
	"huainan":             {32.626389, 116.996944},
	"hyderabad telangana": {17.375278, 78.474444},
	"hyderabad sindh":     {25.39242, 68.373656},
	"ibadan":              {7.387778, 3.896389},
	"indore":              {22.716667, 75.833333},
	"istanbul":            {41.018611, 28.964722},
	"izmir":               {38.412726, 27.138376},
	"jabalpur":            {23.166667, 79.95},
	"jaipur":              {26.916667, 75.816667},
	"jakarta":             {-6.174444, 106.829444},
	"jiddah":              {21.516944, 39.219167},
	"jilin":               {43.850833, 126.560278},
	"jinan":               {36.668333, 116.997222},
	"johannesburg":        {-26.205171, 28.049815},
	"juarez":              {31.733333, -106.483333},
	"kabul":               {34.516667, 69.183333},
	"kaduna":              {10.522239, 7.43828},
	"kalyan":              {19.25, 73.15},
	"kampala":             {0.3155556, 32.5655556},
	"kano":                {11.994354, 8.513807},
	"kanpur":              {26.466667, 80.35},
	"kaohsiung":           {22.6177778, 120.3013889},
	"karachi":             {24.9056, 67.0822},
	"karaj":               {35.8355, 51.0103},
	"kawasaki":            {35.520556, 139.717222},
	"kazan":               {55.78874, 49.122144},
	"khartoum":            {15.5880556, 32.5341667},
	"khulna":              {22.8, 89.55},
	"kiev":                {50.433333, 30.516667},
	"kinshasa":            {-4.3, 15.3},
	"kobe":                {34.691304, 135.182995},
	"kuala lumpur":        {3.166667, 101.7},
	"kumasi":              {6.6833333, -1.6166667},
	"kunming":             {25.038889, 102.718333},
	"lagos":               {6.453056, 3.395833},
	"lahore":              {31.549722, 74.343611},
	"lakhnau":             {26.85, 80.916667},
	"lanzhou":             {36.056389, 103.792222},
	"leon":                {21.116667, -101.666667},
	"lima":                {-12.05, -77.05},
	"london":              {51.514125, -.093689},
	"los angeles":         {34.0522222, -118.2427778},
	"luanda":              {-8.836804, 13.233174},
	"lubumbashi":          {-11.666667, 27.466667},
	"ludhiana":            {30.9, 75.85},
	"luoyang":             {34.683611, 112.453611},
	"lusaka":              {-15.4166667, 28.2833333},
	"madras":              {13.083333, 80.283333},
	"madrid":              {40.408566, -3.69222},
	"maiduguri":           {11.846441, 13.160274},
	"makasar":             {-5.14, 119.4221},
	"managua":             {12.1508333, -86.2683333},
	"manaus":              {-3.113333, -60.025278},
	"mandalay":            {22.0, 96.0833333},
	"manila":              {14.6042, 120.9822},
	"maputo":              {-25.9652778, 32.5891667},
	"maracaibo":           {10.6316667, -71.6405556},
	"mashhad":             {36.297, 59.6062},
	"mecca":               {21.426667, 39.826111},
	"medan":               {3.583333, 98.666667},
	"medellin":            {6.25184, -75.563591},
	"melbourne":           {-37.813938, 144.963425},
	"mexico":              {19.434167, -99.138611},
	"milan":               {45.466667, 9.2},
	"minsk":               {53.9, 27.5666667},
	"mogadishu":           {2.0666667, 45.3666667},
	"monterrey":           {25.666667, -100.316667},
	"montevideo":          {-34.8580556, -56.1708333},
	"montreal":            {45.5, -73.583333},
	"moscow":              {55.752222, 37.615556},
	"multan":              {30.195556, 71.475278},
	"munich":              {48.1500, 11.5833},
	"nagoya":              {35.180198, 136.906739},
	"nagpur":              {21.15, 79.1},
	"nairobi":             {-1.2833333, 36.8166667},
	"nanchang":            {28.55, 115.933333},
	"nanjing":             {32.061667, 118.777778},
	"new delhi":           {28.6, 77.2},
	"new york":            {40.7141667, -74.0063889},
	"nezahualcoyotl":      {19.413611, -99.033056},
	"nizhniy novgorod":    {56.326944, 44.0075},
	"nova iguacu":         {-22.759188, -43.431887},
	"novosibirsk":         {55.0415, 82.9346},
	"odesa":               {46.463934, 30.738551},
	"omsk":                {55, 73.4},
	"ouagadougou":         {12.3702778, -1.5247222},
	"palembang":           {-2.916667, 104.75},
	"paris":               {48.866667, 2.333333},
	"patna":               {25.6, 85.116667},
	"peking":              {39.928889, 116.388333},
	"perth":               {-31.95224, 115.861397},
	"peshawar":            {34.008, 71.578488},
	"philadelphia":        {39.9522222, -75.1641667},
	"phnum penh":          {11.55, 104.9166667},
	"phoenix":             {33.4483333, -112.0733333},
	"pimpri":              {18.616667, 73.8},
	"port harcourt":       {4.777423, 7.013404},
	"port-au-prince":      {18.5391667, -72.335},
	"porto alegre":        {-30.033333, -51.2},
	"prague":              {50.083333, 14.466667},
	"pretoria":            {-25.706944, 28.229444},
	"puebla":              {19.05, -98.2},
	"pune":                {18.533333, 73.866667},
	"qingdao":             {36.098611, 120.371944},
	"qom":                 {34.6401, 50.8764},
	"quito":               {-0.2166667, -78.5},
	"rabat":               {34.013784, -6.844268},
	"rajkot":              {22.3, 70.783333},
	"rangoon":             {16.7833333, 96.1666667},
	"rawalpindi":          {33.6007, 73.0679},
	"recife":              {-8.05, -34.9},
	"rio de janeiro":      {-22.9, -43.233333},
	"riyadh":              {24.653664, 46.71522},
	"rome":                {41.9, 12.483333},
	"rongcheng":           {23.528858, 116.364159},
	"rosario":             {-32.946818, -60.639317},
	"rostov-na-donu":      {47.23135, 39.723284},
	"saint petersburg":    {59.894444, 30.264167},
	"saitama":             {35.895534, 139.67747},
	"salvador":            {-12.983333, -38.516667},
	"samara":              {53.1835, 50.1182},
	"san antonio":         {29.4238889, -98.4933333},
	"san diego":           {32.7152778, -117.1563889},
	"santiago":            {-33.45, -70.666667},
	"santo domingo":       {18.4666667, -69.9},
	"sao paulo":           {-23.473293, -46.665803},
	"semarang":            {-6.9932, 110.4203},
	"seoul":               {37.5985, 126.9783},
	"shanghai":            {31.045556, 121.399722},
	"shenyang":            {41.792222, 123.432778},
	"shenzhen":            {22.533333, 114.133333},
	"shiraz":              {29.6036, 52.5388},
	"singapore":           {1.2930556, 103.8558333},
	"sofia":               {42.6833333, 23.3166667},
	"soweto":              {-26.267812, 27.858492},
	"stockholm":           {59.333333, 18.05},
	"surabaya":            {-7.249167, 112.750833},
	"surat":               {20.966667, 72.9},
	"suzhou":              {31.311389, 120.618056},
	"sydney":              {-33.861481, 151.205475},
	"tabriz":              {38.08, 46.2919},
	"taichung":            {24.1433333, 120.6813889},
	"taipei":              {25.0391667, 121.525},
	"taiyuan":             {37.726944, 112.470833},
	"tangerang":           {-6.178056, 106.63},
	"tangshan":            {37.3325, 114.701389},
	"tashkent":            {41.3166667, 69.25},
	"tbilisi":             {41.725, 44.7908333},
	"thana":               {19.2, 72.966667},
	"tianjin":             {39.142222, 117.176667},
	"tijuana":             {32.533333, -117.016667},
	"tokyo":               {35.685, 139.751389},
	"toronto":             {43.666667, -79.416667},
	"tripoli":             {32.8925, 13.18},
	"ufa":                 {54.785167, 56.045621},
	"umm durman":          {15.6361111, 32.4372222},
	"urumqi":              {43.800965, 87.600459},
	"vadodara":            {22.3, 73.2},
	"valencia":            {10.1805556, -68.0038889},
	"vancouver":           {49.25, -123.133333},
	"varanasi":            {25.333333, 83},
	"vienna":              {48.2, 16.366667},
	"visakhapatnam":       {17.7, 83.3},
	"volgograd":           {48.71939, 44.501835},
	"warsaw":              {52.25, 21},
	"wuhan":               {30.580125, 114.273405},
	"wuxi":                {31.568873, 120.288573},
	"xian":                {34.258333, 108.928611},
	"xianyang":            {34.33778, 108.70261},
	"xinyang":             {32.095833, 114.120278},
	"xuzhou":              {34.180454, 117.15707},
	"yaounde":             {3.8666667, 11.5166667},
	"yekaterinburg":       {56.8519, 60.6122},
	"yerevan":             {40.1811111, 44.5136111},
}

/* The list has been extracted with the following awk script:
# Country,City,AccentCity,Region,Population,Latitude,Longitude
BEGIN { FS="," }
{
        if (($5 > 1000000) && (NR > 1)) {
                print "\"" $2 "\":{" $6 "," $7 "},"
        }
}
*/
