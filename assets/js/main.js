var seq = 0,
    delays = 80,
    durations = 500;

function startAnimationForLineChart(chart) {

    chart.on('draw', function (data) {
        if (data.type === 'line' || data.type === 'area') {
            data.element.animate({
                d: {
                    begin: 600,
                    dur: 700,
                    from: data.path.clone().scale(1, 0).translate(0, data.chartRect.height()).stringify(),
                    to: data.path.clone().stringify(),
                    easing: Chartist.Svg.Easing.easeOutQuint
                }
            });
        } else if (data.type === 'point') {
            seq++;
            data.element.animate({
                opacity: {
                    begin: seq * delays,
                    dur: durations,
                    from: 0,
                    to: 1,
                    easing: 'ease'
                }
            });
        }
    });

    seq = 0;
}

function init(result) {
    /* ----------==========     Daily Sales Chart initialization    ==========---------- */
    // console.log(result.historyRate);
    const arrRate = result.historyRate.map(v => v.exchangeRate * 100);
    const arrDate = result.historyRate.map(v => v.effectiveDate);
    // console.log(arrRate);
    // console.log(arrDate);

    dataDailySalesChart = {
        labels: arrDate.reverse(),
        series: [arrRate.reverse()]
    };

    const down = Math.floor(Math.min(arrRate));
    const up = Math.ceil(Math.max(arrRate));

    optionsDailySalesChart = {
        lineSmooth: Chartist.Interpolation.cardinal({
            tension: 0
        }),
        low: down,
        high: up, // creative tim: we recommend you to set the high sa the biggest value + something for a better look
        chartPadding: {
            top: 0,
            right: 0,
            bottom: 0,
            left: 0
        },
    }

    var dailySalesChart = new Chartist.Line('#dailySalesChart', dataDailySalesChart, optionsDailySalesChart);

    startAnimationForLineChart(dailySalesChart);
}

