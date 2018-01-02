package statistica_activity

//
//#include <string.h> 
//
//#define STRING_EQUAL(a,b) (_stricmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (_stricmp(a,b)!=0)
//#define STRING_SET(a,b)   strcpy(a,b)
//
//
//void _BTrees_Prediction_T5_38_41(
//      const double * _ozone__,
//      const double * _humidity__,
//      const double * _dewPoint__,
//      const double * _pressure__,
//      const double * _windSpeed__,
//      const double * _value__,
//      const double * _Index__,
//      const char * _pollutant__,
//      char * _pRet
//   )
//   {
//     double _MaxValue;
//     int _i;
//     double _den;
//     double _LearningRate;
//     double _PredictProb[2];
//     _MaxValue = -1.0E30;
//     _den = 0;
//     _LearningRate = 0.100000;
//
//     for(_i=0;_i<2;_i++)
//     {
//         _PredictProb[_i] = 0;
//     }
//
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += 1.000000 * -6.66913671730271e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += 1.000000 * 1.58701602763874e-001;
//
//    }
//    else {
//    _PredictProb[0]  += 1.000000 * -4.99444444444444e-001;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.43533676598930e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -6.83716151904799e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.71276658509456e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.41290662723322e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.21069486504614e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.76533167435336e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.34844942200266e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 9.88669624441755e-004;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.91767966677910e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.30746101093808e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.68598201721800e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.91943261238780e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.28750772285108e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -3.00533780718418e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.78149359061554e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.26078457717503e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -2.34515821834571e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.60321657814334e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.24452644398622e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -6.99244327740577e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.08275278784428e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.23466800930454e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.91225648370428e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.79925665043033e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.22116525162483e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.44265009235865e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.01345576909095e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.19990184926426e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -3.12297150705546e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.28258878813168e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.17964848136723e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 4.88145131379156e-004;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.40502739052933e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.16996751972666e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 5.65587092657889e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.58544752735972e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -5.02326124589027e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.08155299585091e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.10765874281867e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.14923686516817e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.54816232597826e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.01030399884284e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.13159844500170e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 9.86165857389699e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.75630411437851e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.85065477230875e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.21781848473017e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.55900874123942e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.24995227822550e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.04839188671409e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.41379910914591e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.29432172653339e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.62722860137992e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.41379973189204e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -7.80553024599372e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.11190952428327e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.72997259575475e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -4.24129214598838e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.20603384469481e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.04214089408641e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.07095562752523e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.21139742567601e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.31626763985876e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.77588260286917e-004;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.04219134832464e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.39671757033541e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.11350477095878e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 6.71239491067529e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.13605730814760e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.43915268203897e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.02446711411868e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.49805156757043e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -2.80226117088387e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 3.35589711428697e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.50390870942274e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.10638765249783e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 7.61157608177155e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.97522603787012e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -4.46276069532255e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 3.51005971109850e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.47895424308354e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.38183133949738e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.09639968606267e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.54978180447181e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.35182688986272e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.49956285603456e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.83617847106106e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.65594290749769e-003;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -7.24781268563297e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.74880039845349e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.09211874939280e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.04447035345970e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.77576327983786e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -4.86343464682038e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 3.39968389980685e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.14514376438467e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.67938741201842e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.05729161904198e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.10285563956002e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.65796696966619e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -4.54223300599861e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.23160198191589e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.76179540627094e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.73366417354210e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.25868532716186e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 3.88432604210403e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.25992399000099e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -8.77816525117369e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.11939073011537e-001;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.30212191852629e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.61842931764624e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.07860571514240e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.43633933748420e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.11801762756374e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -9.73833112422494e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.16807413831276e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.12087833522420e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.18542518119845e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.13165677677904e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.15930997609061e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -9.45556444141354e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.18250663656384e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.51165264621336e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.00000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.08253915445628e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.00000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.43220790106957e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.32710431770259e-004;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 2.42409738451086e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -2.18058464886713e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.91544189210106e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -5.04331405610929e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.99846605534107e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.44769297870484e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 2.46366254514046e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -3.18718248846993e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.47710122075474e-004;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.79500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.19463375103708e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.79500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.53334991794186e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.16386033219161e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.06865161065799e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 4.91195529067842e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.06534586674729e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.86060933251467e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 2.25516616668772e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.60304687176006e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.29862055917490e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 7.00213411280753e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.42249110629193e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.05259221374021e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.12132151379869e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.41790868127104e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.12928656353405e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -8.23612978399469e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.59103502581614e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 2.23813798297490e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -3.47822553709377e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.33038415251032e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.60168897615949e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -3.68547405676478e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.00807048261037e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 4.07656593113709e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -3.39488016288110e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.53933998755768e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.21612601002145e-001;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.38911626304440e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.38218059456748e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -7.39506681314395e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.32895263117508e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.69290712121121e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.00000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.64006294038946e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.00000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -9.02926978097944e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.92349147653127e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.62622098139782e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.93855352828441e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.96691342285364e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.36367057387583e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.47975082917858e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.91902443326711e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.82711314859384e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.42192774336860e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.21931917474971e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -4.25021102309123e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 3.45478910607206e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.35115756377370e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.37670334020403e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.11862729273022e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.89208012324510e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.30441245627952e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.53675704284707e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.44085177415571e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.46067996614992e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.56191414431042e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.62496072080712e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.88378893674508e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.74483620387354e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.82316364721388e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.27345207096055e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.25422352727840e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.59925302294342e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.22568169342267e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.39023094077787e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.47438999052512e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -4.28414872999872e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.95077590117710e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.71502510396777e-004;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.62252955266148e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.23703366941372e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -8.77917740409668e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.76880531795934e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.57666361868797e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.81840638549054e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.09675292505589e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -8.44838260978611e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.67540171469183e-003;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 1.21312395264626e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -2.87998956462548e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.44012884566749e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.14101660588492e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.44269967434424e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.29454718508235e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.19088028569221e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.53832630275037e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.60246219231889e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.10478799609207e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.61085706096745e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.76797394015758e-003;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 7.06724208235010e-003;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -3.85347846801714e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.26408975698113e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 2.49126019441748e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -4.61706688971833e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.20084319391975e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.40586801579097e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.66926132186785e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.98915788293497e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 2.04039850217047e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -3.78838603637487e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.85888008974691e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -2.12065624242333e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 7.42066969097006e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.95769365091650e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.05799633056740e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 7.03064667203308e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.11022329052471e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.76772377484563e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 2.69738091467020e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.41181162876995e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -5.64931936351116e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 2.17774026524829e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.39876362444956e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.21571987942176e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.22780795402591e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.49605436581907e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.06347980170237e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.95478913905480e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.19675730913709e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 2.22998172572441e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -4.92272270410149e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.02864590252236e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.18610106024875e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -4.42862332241145e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.77078350694881e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 2.33214954491474e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -3.25305463889034e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.41806870928136e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.14110317657294e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.43565665271077e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.73261600356431e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.13246859911683e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -8.58283130108468e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.08172610802622e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.15710698886492e-001;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.28871979903325e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.69849858779676e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.23133033053883e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.35286942076776e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.67578313143521e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.15853514350661e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.00146875899603e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.02074818089118e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.58205821621670e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.33992085408285e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.14898461380066e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.05274005535323e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 7.27131118426977e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.54276531256137e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.05167583896736e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 8.32802604164687e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.58520833356768e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.15383439302827e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.27767698518819e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.25986290755911e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.98705502795463e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.14785046663290e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.54897334517723e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.53300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -7.05218323115237e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.53300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.96558947599691e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.85892830532684e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.69539468709675e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.42525309442642e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.17987186932592e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.88294883516173e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.09071352823934e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.08111935445300e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.89635182845500e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.54874544468219e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 8.26592525959586e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.10000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.73927315413438e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.10000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.39606122694492e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.67043248153249e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -2.24988281688581e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 5.47075792052514e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.67426437696870e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.53300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -4.35996385183471e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.53300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.10799534144901e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.24689177956736e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.83345000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 3.93586266247081e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.83345000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.21111360486107e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.78125769525548e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.71500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -4.94296671945024e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.71500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.37193910786635e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.42621306096744e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01621500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.96583232199731e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01621500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 4.10428223284074e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.48796879202024e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.09236120362905e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.76650397579360e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.42009430468948e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.04880406464037e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 9.77536865178962e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.23091806653738e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.58460046278170e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.48280195828440e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.32262785954862e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.71513139025815e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.67149150437192e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.08741794706131e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.03835567462426e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.46151989141623e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.77414578448212e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.62820416536150e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -4.26509265062216e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.81607312783452e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.18531061323402e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.74273234270473e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.20384110822293e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.19783800225444e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 6.09362717401791e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.85765619127433e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.97364538055461e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.80469514170815e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.16412994139430e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.35642956574773e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.13585244797419e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.50487230779882e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.85370187976144e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 1.06295894522726e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.50276677715633e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.84935660350743e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 7.29229651268349e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -3.46954526818524e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 7.33835667482187e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -4.73047985306441e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.85062514334033e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.83345000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -4.67408831873253e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.83345000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 1.35373839948699e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.49716124396114e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.57184867569717e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.60424572719010e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.59954775187222e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -1.87470328656635e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 3.54452253727809e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.40756290096024e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.64350000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.76076498154085e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.64350000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.24244929402372e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.73239530614560e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.80519762559255e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 3.81552186075685e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.84592957097346e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -5.03629813260487e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.43148958795106e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.14609560250799e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.34016852416754e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 7.16900590539712e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.08966287847737e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -5.04314367129879e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 8.14169422119856e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.01190460932841e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.85185000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -3.51369763749694e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.85185000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 6.29590715727148e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 5.19458556783841e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -1.79293748430955e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.23587503936322e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.34364543622266e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.49557385954797e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 5.47132349345115e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.40763762199083e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.51599185123417e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.90943007460704e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.60041564789738e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.48354546677807e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -9.70652252408498e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.45736089873226e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 9.24267752742212e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.07361736478862e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -4.47226299863373e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -5.08982625948053e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.46129283694890e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.32734624489319e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.83272158067163e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.07784602102289e-001;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.42700848728211e-004;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.99210349469814e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.40130360244607e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.47922799227552e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.93436369777913e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.75917170296564e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 4.16876721972059e-004;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 3.66870527145254e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.37263317315312e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.51342856369726e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.91218703576674e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.33133092332009e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.02898349397602e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.32094749637625e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.61338975510636e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 7.01282089667092e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 4.36711895685884e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.49100000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -1.35068226946739e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.38572036262775e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[0]  += _LearningRate * -3.38865479095197e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[0]  += _LearningRate * 2.61059820396997e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.41549036387880e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -1.94034273594232e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 2.34661061636645e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.20601926133396e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * -1.94595728601943e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[0]  += _LearningRate * 5.57228943813157e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.28765931113168e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.32403689553570e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.78162151787043e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -5.39073138848789e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 2.27034436477025e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.83299156148605e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 6.27063796901431e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -5.08277269605400e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 2.45857835812554e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.05425803729581e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[0]  += _LearningRate * -3.20585510042772e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[0]  += _LearningRate * 3.68094257671220e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.29917225921223e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 5.94792555556106e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -3.15419741721137e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.24462693649483e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.54509416763539e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.44947801214446e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.81256219342237e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -8.08959360721277e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.09866401584246e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.12423376107941e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -3.32324719153120e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 6.12762545143062e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 1.94942943941643e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -7.82143631121142e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 1.16251506085728e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.24737685859759e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 2.41391454168878e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -4.11749782786012e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.27651521775610e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -5.36407030299267e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.88961160567964e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.09892626093835e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 2.30992241915433e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -5.00410788433945e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 9.32011374708855e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -7.97345848569882e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 7.40716692388498e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -7.78714895718359e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 4.71969083881945e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.17858716060839e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.65545964206335e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -6.46473309491463e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 8.24894109129228e-003;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 2.23608965820344e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * 3.45142806551736e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[0]  += _LearningRate * -4.33437861040792e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * 3.72819841501722e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.09113512107096e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -5.47180464920803e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -9.71638117697895e-004;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[0]  += _LearningRate * -3.42705306748576e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[0]  += _LearningRate * 3.73658467118700e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.81067075971562e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * 3.99196095971953e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[0]  += _LearningRate * -2.65238977377562e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -1.31398175949160e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * 1.29645477697007e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[0]  += _LearningRate * -2.32269628417009e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -6.70513064503997e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * -2.99292447849604e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[0]  += _LearningRate * 4.52308110922565e-002;
//
//    }
//    else {
//    _PredictProb[0]  += _LearningRate * -2.37991911440213e-002;
//
//    }
//
//   if(_MaxValue <= _PredictProb[0])
//   {
//     _MaxValue = _PredictProb[0];
//     STRING_SET(_pRet,"0");
//   }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += 1.000000 * 6.56456346294938e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += 1.000000 * -1.56504644497305e-001;
//
//    }
//    else {
//    _PredictProb[1]  += 1.000000 * 3.18245657605554e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.44997729534278e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 3.89964345938229e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.08821264374435e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.40787921081929e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.99766856574010e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.03482354750383e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.34586813144086e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.34419330987252e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.34723061743789e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.32340705324529e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 1.36712575882423e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.30769741412672e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.28226502070750e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 1.23143801339795e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.80642705319828e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.53505292899148e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.77552451780782e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.76513684622663e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.09554541891122e-001;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.01749494278874e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.11396502363888e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.22814994282671e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 6.48244224873076e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.28282601506326e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.76429333943318e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.89137298908726e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.55380952215946e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.37141828914481e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.66602302308251e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.16934543488182e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.18373614601476e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 1.57621117136089e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.50095963492938e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 8.58927343494018e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.25924405432142e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.88561027164965e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.15899277694855e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -6.56867264887488e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.86766875847449e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.40098598764451e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.96154393783276e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.03142586824997e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.14174486710231e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -3.47135839472212e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.26873019322671e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.76538683759090e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.80107219355911e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.93110232042359e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.13370991952259e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -4.60068504482152e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.24643680387300e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.02202954315855e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.59469173151328e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.21777757171399e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.40599825678378e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.62273682857907e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.10419319802624e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.12517756695720e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -8.19480631232717e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.04180128643326e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.51231298684787e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.43073997059549e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.20467246980070e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 4.14751295919003e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -6.52556689703885e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.94837255150675e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 4.80463689798196e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -2.67898248624907e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.27824450109555e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01550000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.01867581557382e-001;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01550000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -1.65508797604634e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.40505330588168e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.00000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.98617328077404e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.00000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.92950593236812e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.46566662771397e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.14607558814655e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.02317387133932e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.14499207155176e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 8.85488152787314e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.05301424267633e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.82793359754238e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.09964030046869e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.11295531850893e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.18465471050022e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.09693606670517e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -5.20677522178077e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.93965438259206e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.70926667523097e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 8.40688566849577e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.62695285090644e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.25391098615086e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.10625856326183e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.88189871220529e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.09012375147895e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.19794798721517e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.14163905544721e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 8.16397162806920e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.27196416020490e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.52085086463017e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.93805378401546e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.17391231236292e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.94871305036286e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.06093530630667e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -8.92144669568216e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.40883515881949e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.56895590052456e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.66176920067912e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.62084433417516e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.08003422573851e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -3.29614838876520e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.26032725482485e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 8.42528321002910e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.82022941226940e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.13146930279235e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.26220060882417e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.99908883335091e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.07724222294006e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.07381377026411e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -7.75732297833151e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.13856116119391e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.67088085076053e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -9.79008648733719e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.25633398761970e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 9.84890960583584e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.80968287274733e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.81733307820669e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 3.16000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.22103185422625e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 3.16000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.63696235860833e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.55881963367579e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.05852806548928e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.20179232294323e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.49376497485988e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.07266253954149e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -7.07571565574105e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.49291341200064e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.60981443485499e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.54558553620984e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.26659671661712e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.82206800121818e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.59826980048121e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.60177487078761e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -4.05390278487649e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 4.89466187615113e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.88561369367635e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.19780651953984e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.95471955606762e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.85586958459473e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 2.23608310861151e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -6.52931373014582e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.21983026144079e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 2.06941582793706e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -9.24581346392627e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.09031867989301e-004;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.42312984287393e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.18611282777620e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.45634184332626e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.00815901505163e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.88739540545553e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.94113036687549e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.07471058078916e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.29441493132600e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.99313218772462e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.75474308534206e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 3.36068332155563e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.83031924952511e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.93293985993705e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.43531736620388e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.60125245787821e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 2.49042403251787e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -3.37652137581229e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.76762443472137e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.20659066394961e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.77953806672564e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.86614262624558e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.71500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 7.27106515864267e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.71500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.35315372069198e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.41923956348463e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.87187113627796e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.75123384290793e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.39074812513179e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * -1.82422596740089e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * 2.57806683215664e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.04904614839319e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * -3.29876223085384e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * 2.35104448927778e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.90576222586034e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.69432160174894e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.22714195105497e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.54780346045035e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.06133192491811e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -3.28382719104841e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.26818502949234e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.65022662368994e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.96136268692236e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.79779721875993e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.87673390083628e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.07489307418245e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.35634825208694e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.51139168146760e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.33102449434222e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.84090359059590e-003;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * -4.46977543419266e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * 5.01377342944659e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.51609666740054e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.96714539229997e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.55621725708257e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.00774458012554e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.34946002785312e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.99419749553802e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.78063436255006e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.69844309761196e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.86841056952078e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.04500648154394e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.63197996426507e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.71804006494050e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.12369802645354e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 3.39515516340161e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -3.99118391446751e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.86717778503093e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.27618598619393e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -5.96454127595364e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.34931995693254e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.59440329426939e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.95715127110216e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.75270915141047e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.00000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.33051093237168e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.00000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.16781017358162e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.57227026013379e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.93409630895575e-003;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.60491801910401e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.86183396903389e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.63297592865994e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.91330252582198e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.92846592837644e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 2.86442934208885e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -4.36074651525647e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.37943471854127e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.88087513767859e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.89628990836715e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.33357079265019e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 8.54513371635864e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -9.65752759014230e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.13656346909452e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.51506012634942e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 6.35000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.09325497870510e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.42483257961844e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 2.95214673293994e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01614500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -2.92286396741552e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.00498096533675e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.78061670597758e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.49480343848328e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.47542186115836e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 1.09935721967541e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.79360000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -2.06799445616469e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.52327717151181e-003;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 1.95798766271196e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -5.07405625641866e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.45039762257016e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -5.44955788280980e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.08853673565161e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.54446148463970e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.36919387616822e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.04530318687631e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.52690352159124e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.39384987005884e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.90840729264519e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.50868815621588e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -2.26049668005135e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 2.25058214886434e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.62553273408752e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.83142596352861e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 7.07009259733346e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.60522426176450e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.49415281455307e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.43587605158360e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.36261637136395e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.86396125651748e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.33137474605784e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.61772987637932e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.40355656911702e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.92200367522041e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.51514385718510e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.58366622928774e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.41577000196754e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 9.77953009527180e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 7.23817227918369e-003;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -7.95050502212877e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.57436243481653e-003;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * -4.04889811215146e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * 2.27299034980440e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.39007498223868e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.64799253573923e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.31008571291661e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.02562048245368e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 7.69488189166077e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.06044971287868e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.36217222131191e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.56367874003492e-003;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.02782387619665e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.25264454957729e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.87758505135691e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.25530104115183e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.88880414895690e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.59059350769337e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -8.94422761559864e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.17279518914915e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.65391261836929e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.84947998039617e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.41010156840641e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.46155452040059e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.43126330570953e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.50336613822349e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.91888715349514e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.02364990762553e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 5.05507523813632e-004;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.04792110327949e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -6.83865145525480e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.05512305319741e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.03073866720903e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.38300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 7.61612669921522e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.99744997163065e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.31371088398936e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 9.27102254285729e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.83109912652435e-003;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -2.28459176795536e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 7.16872734860314e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.10529940010890e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.24355655804133e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.12439624136068e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.05556829244605e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -9.04343012230374e-003;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 3.84433871812160e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.02306618818528e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.47052776842737e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.45804694377044e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.57900718951075e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.53300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.11467972621269e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.53300000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.05599168171610e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.11726324703216e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.41390494319217e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01629000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -8.31589432992651e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.80034449379991e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 6.30403013536007e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -7.32933366011406e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.97402325386919e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.03381121593663e-001;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.82167815231254e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.38143384274055e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.64842110619968e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.43579028425162e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.26506587630308e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.04581404358768e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -8.98048484508442e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.54626090873197e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.94406229415113e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.13102732715569e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.85361341079654e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 1.42461429845630e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -3.91722331546554e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -8.75381464695547e-003;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -1.69316934698269e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 5.18156844799564e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.43648437728557e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.10000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.62675368335443e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.10000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.01934348935496e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.03003204902956e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 3.16200852541165e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.77541201622611e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.64439647875592e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.05438950474443e-001;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.86558917830063e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -6.10397592308424e-003;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.37099045319587e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.79500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.50454393863880e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.46052130172928e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.77448163566787e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 6.21874556274290e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 7.56667548412848e-004;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 5.04603124022475e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.88471007370081e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.01183958465231e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -2.16446546977765e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.67856910812341e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.12626203624827e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.69098791558782e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.04992035975884e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.42520759279485e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.59505466974226e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.68150000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.43485103797976e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 8.91320034029552e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.63744618862560e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.69900000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.30701054757897e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.67496129261536e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * -2.39388327427622e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * 2.24427488647769e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.02803117691592e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 3.22500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.46023468497272e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 3.22500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.57744827188551e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.39720793804263e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.62500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.84056514815200e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.62500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.27537039018669e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.85401246791007e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.09045393924023e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.65000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -4.64098324321489e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.03954091407559e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 3.19783085271610e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -1.35442386516867e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.29288421604709e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.48474376524997e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.07612231378030e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.17805716820646e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.00182819404863e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 4.75000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.12997735916922e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.09335234297855e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.91312185084551e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -4.09454766706951e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.09195802695683e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.27335564078418e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.25132910484400e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.31203327714060e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -5.29483929555251e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.55000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.62673572820374e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.20386356836191e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.90270907181832e-003;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.56395872268627e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.74250058213231e-002;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 3.36861776861580e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.64350000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.98364671948349e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.96822341453989e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 7.21234835137601e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 1.70000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -8.29517844230455e-003;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.49495819947904e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.40373526343783e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.08977848159485e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.60358771764735e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.13451670206239e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.58551260821149e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.08542248197897e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.53653944868076e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -5.64520850349940e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -9.95756448552941e-003;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.69874633948948e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.52243069199926e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 4.37235707067519e-004;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -6.06196402849398e-003;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.79076920990459e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.19453709261965e-002;
//
//    }
//    if ( _humidity__ != NULL && *_humidity__ <= 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * 2.21596785105763e-002;
//
//    }
//    else if ( _humidity__ != NULL && *_humidity__ > 7.60000000000000e-001 ) {
//        _PredictProb[1]  += _LearningRate * -4.66599000250782e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.71345486170870e-002;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * -2.78267242373816e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * 4.05047232048040e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.19189796230128e-003;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -3.16579369189255e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.41648148678199e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 6.05748011751671e-003;
//
//    }
//    if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Particles (PM2.5)")  ) {
//        _PredictProb[1]  += _LearningRate * -3.92379094443900e-002;
//
//    }
//    else if ( _pollutant__ != NULL &&  STRING_EQUAL(_pollutant__,"Ozone")  ) {
//        _PredictProb[1]  += _LearningRate * 2.91937426720337e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.41304177199917e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 3.22500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.61637786288252e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 3.22500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -3.26459504724078e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 2.25974180141007e-002;
//
//    }
//    if ( _windSpeed__ != NULL && *_windSpeed__ <= 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 5.30996203678538e-002;
//
//    }
//    else if ( _windSpeed__ != NULL && *_windSpeed__ > 1.51500000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -1.05429384527903e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.76701343600335e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.55157446107361e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01621500000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -5.74642402115201e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -3.87656972025981e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 2.81283339551167e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -4.03431489111555e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.08729910578691e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -7.27096690167925e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 2.85000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 2.31791441444014e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -4.01609950745606e-003;
//
//    }
//    if ( _dewPoint__ != NULL && *_dewPoint__ <= 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.66833237328308e-002;
//
//    }
//    else if ( _dewPoint__ != NULL && *_dewPoint__ > 6.72600000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -6.24448935881059e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.06543809821514e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 1.01395774519182e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 7.40000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.09067790195926e-001;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -7.13837071284122e-003;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -2.56792427753995e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 1.17402631614914e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.17602846111337e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.51115899651669e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.65000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 1.06462187934115e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 3.62886090990873e-002;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -2.58343941261438e-002;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.84370000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 2.17938010101939e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -5.75275258620346e-003;
//
//    }
//    if ( _ozone__ != NULL && *_ozone__ <= 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * 5.03913376530802e-001;
//
//    }
//    else if ( _ozone__ != NULL && *_ozone__ > 2.70335000000000e+002 ) {
//        _PredictProb[1]  += _LearningRate * -1.11903979208441e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.18734538235490e-002;
//
//    }
//    if ( _pressure__ != NULL && *_pressure__ <= 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * -2.05800124096242e-002;
//
//    }
//    else if ( _pressure__ != NULL && *_pressure__ > 1.01610000000000e+003 ) {
//        _PredictProb[1]  += _LearningRate * 2.21824298091052e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -1.49896119450956e-002;
//
//    }
//    if ( _Index__ != NULL && *_Index__ <= 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * -2.59203029636837e-002;
//
//    }
//    else if ( _Index__ != NULL && *_Index__ > 7.80000000000000e+000 ) {
//        _PredictProb[1]  += _LearningRate * 4.34702078600255e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * -2.15661756154579e-002;
//
//    }
//    if ( _value__ != NULL && *_value__ <= 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * 4.05980058078971e-002;
//
//    }
//    else if ( _value__ != NULL && *_value__ > 5.50000000000000e+001 ) {
//        _PredictProb[1]  += _LearningRate * -1.89467337743724e-002;
//
//    }
//    else {
//    _PredictProb[1]  += _LearningRate * 1.64406404094335e-002;
//
//    }
//
//   if(_MaxValue <= _PredictProb[1])
//   {
//     _MaxValue = _PredictProb[1];
//     STRING_SET(_pRet,"1");
//   }
//
//}
//
import "C"

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log Is the Default package logger
var log = logger.GetLogger("activity-tibco-inference")

// InferfenceActivity is an Activity that is used to Invoke a ML Model using flogo-ml framework
type ModelActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a New InferfenceActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ModelActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ModelActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *ModelActivity) Eval(context activity.Context) (done bool, err error) {

	var i1 = C.double(float64(context.GetInput("_ozone__").(float64)))
	var i2 = C.double(float64(context.GetInput("_humidity__").(float64)))
	var i3 = C.double(float64(context.GetInput("_dewPoint__").(float64)))
	var i4 = C.double(float64(context.GetInput("_pressure__").(float64)))
	var i5 = C.double(float64(context.GetInput("_windSpeed__").(float64)))
	var i6 = C.double(float64(context.GetInput("_value__").(float64)))
	var i7 = C.double(float64(context.GetInput("_Index__").(float64)))
	var i8 = C.CString(context.GetInput("_pollutant__").(string))

	var result C.char
	C._BTrees_Prediction_T5_38_41(&i1, &i2, &i3, &i4, &i5, &i6, &i7, i8, &result)

	context.SetOutput("result", C.GoString(&result))

	return true, nil
}

