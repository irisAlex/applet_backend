<template>
    <div class="lineCharts-box">
        <div ref="echart" class="lineCharts-box-echarts" :style="`width : ${chart?.clientWidth}px`" />
    </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue'

const chart = ref(null)
const echart = ref(null)
const initChart = () => {
    chart.value = echarts.init(echart.value)
    setOptions()
    document.addEventListener('resize', () => {
        chart.value?.resize()
    })
}
const xLabel = ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
const reworkData = [2, 3, 1, 7, 7, 8, 2, 7, 4, 5, 7, 4]
const repairData = [9, 8, 4, 6, 7, 9, 3, 4, 5, 1, 1, 4]
const partsData = [1, 3, 4, 3, 4, 4, 4, 4, 5, 8, 3, 2]
const passData = [7, 2, 9, 5, 6, 9, 2, 4, 5, 9, 1, 3]
const setOptions = () => {
    chart.value.setOption({
        backgroundColor: 'transparent',
        tooltip: {
            trigger: 'axis',
            backgroundColor: 'transparent',
        },
        legend: {
            align: 'left',
            right: '5%',
            top: '0%',
            type: 'plain',
            textStyle: {
                color: '#75BFA5',
            },
            itemGap: 25,
            itemWidth: 20,
            icon: 'path://M0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z',
            data: [
                {
                    name: '返工',
                },
                {
                    name: '返修',
                },
                {
                    name: '配做',
                },
                {
                    name: '让步放行',
                },
            ],
        },
        grid: {
            top: '15%',
            left: '4%',
            right: '2%',
            bottom: '8%',
            // containLabel: true
        },
        xAxis: [
            {
                type: 'category',
                boundaryGap: false,
                axisLine: {
                    // 坐标轴轴线相关设置。数学上的x轴
                    show: false,
                    lineStyle: {
                        color: '#e1e1e1',
                    },
                },
                axisLabel: {
                    // 坐标轴刻度标签的相关设置
                    textStyle: {
                        color: '#92969E',
                    },
                    formatter: function (data) {
                        return data
                    },
                },
                splitLine: {
                    show: false,
                    lineStyle: {
                        color: '#192a44',
                    },
                },
                axisTick: {
                    show: false,
                },
                data: xLabel,
            },
        ],
        yAxis: [
            {
                name: '单位：单',
                nameTextStyle: {
                    color: '#777',
                },
                min: 0,
                splitLine: {
                    show: true,
                    lineStyle: {
                        color: '#e1e1e1',
                    },
                },
                axisLine: {
                    show: false,
                },
                axisLabel: {
                    show: true,
                    textStyle: {
                        color: '#92969E',
                    },
                    formatter: function (value) {
                        if (value !== 0) {
                            return `${value}`
                        }
                        return value
                    },
                },
                axisTick: {
                    show: false,
                },
            },
        ],
        series: [
            {
                name: '返工',
                type: 'line',
                symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
                showAllSymbol: true,
                symbolSize: 0,
                smooth: true,
                lineStyle: {
                    normal: {
                        width: 0,
                        color: '#596fc8', // 线条颜色
                    },
                },
                areaStyle: {
                    // 区域填充样式
                    normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                            {
                                offset: 0,
                                color: '#596fc8'
                            },
                            {
                                offset: 1,
                                color: '#596fc8'
                            }
                        ], false),
                        shadowColor: 'rgba(102,102,255,0.52)', // 阴影颜色
                        shadowBlur: 0,
                    },
                },
                data: reworkData,
            },
            {
                name: '返修',
                type: 'line',
                symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
                showAllSymbol: true,
                symbolSize: 0,
                smooth: true,
                lineStyle: {
                    normal: {
                        width: 0,
                        color: '#abd197', // 线条颜色
                    },
                },
                areaStyle: {
                    // 区域填充样式
                    normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                            {
                                offset: 0,
                                color: '#abd197'
                            },
                            {
                                offset: 1,
                                color: '#abd197'
                            }
                        ], false),
                        shadowColor: 'rgba(102,204,0,0.52)', // 阴影颜色
                        shadowBlur: 0,
                    },
                },
                data: repairData,
            },
            {
                name: '配做',
                type: 'line',
                symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
                showAllSymbol: true,
                symbolSize: 0,
                smooth: true,
                lineStyle: {
                    normal: {
                        width: 0,
                        color: '#fdc75c', // 线条颜色
                    },
                },
                areaStyle: {
                    // 区域填充样式
                    normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                            {
                                offset: 0,
                                color: '#fdc75c'
                            },
                            {
                                offset: 1,
                                color: '#fdc75c'
                            }
                        ], false),
                        shadowColor: 'rgba(102,204,0,0.52)', // 阴影颜色
                        shadowBlur: 0,
                    },
                },
                data: partsData,
            },
            {
                name: '让步放行',
                type: 'line',
                symbol: 'circle', // 默认是空心圆（中间是白色的），改成实心圆
                showAllSymbol: true,
                symbolSize: 0,
                smooth: true,
                lineStyle: {
                    normal: {
                        width: 0,
                        color: '#ed6864', // 线条颜色
                    },
                },
                areaStyle: {
                    // 区域填充样式
                    normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                            {
                                offset: 0,
                                color: '#ed6864'
                            },
                            {
                                offset: 1,
                                color: '#ed6864'
                            }
                        ], false),
                        shadowColor: 'rgba(117,191,165,0.52)', // 阴影颜色
                        shadowBlur: 0,
                    },
                },
                data: passData,
            },
        ],
    })
}

onMounted(() => {
    nextTick(() => {
        setTimeout(() => {
            initChart()
        }, 300)
    })
})

onUnmounted(() => {
    if (!chart.value) {
        return
    }
    chart.value.dispose()
    chart.value = null
})
</script>
<style lang="scss" scoped>
.lineCharts-box {
    height: 360px;
    overflow: hidden;
    position: relative;

    &-echarts {
        position: absolute;
        bottom: 0;
        left: 0;
        right: 0;
        z-index: 2;
        width: 100%;
        height: 100%;
    }
}

.in-line {
    --color: #5BC2A4;
}

.out-line {
    --color: #DF534E;
}
</style>