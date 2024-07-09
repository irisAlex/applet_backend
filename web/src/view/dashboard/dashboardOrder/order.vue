<template>
    <div class="ReclaimMileage-box">
        <div ref="echart" class="ReclaimMileage-box-echarts" :style="`width : ${chart?.clientWidth}px`" />

        <div class="ReclaimMileage-box-data">
            <div class="ReclaimMileage-box-data-center" style="margin-right: 50px;">
                <div>报废</div>
                <div class="number">{{ rework.length }} </div>
                <div>昨日:{{ rework.length }}</div>
            </div>
            <div class="ReclaimMileage-box-data-center" style="margin-right: 50px;">
                <div>返工</div>
                <div class="number">{{ rework.length }} </div>
                <div>昨日:{{ rework.length }}</div>
            </div>
            <div class="ReclaimMileage-box-data-center" style="margin-right: 40px;">
                <div>返修</div>
                <div class="number">{{ repair.length }} </div>
                <div>昨日:{{ repair.length }}</div>
            </div>
            <div class="ReclaimMileage-box-data-center">
                <div>让步</div>
                <div class="number">{{ pass.length }} </div>
                <div>昨日:{{ pass.length }}</div>
            </div>
            <div class="ReclaimMileage-box-data-center" style="margin-left: 30px;">
                <div>配做</div>
                <div class="number">{{ parts.length }}</div>
                <div>昨日:{{ parts.length }}</div>
            </div>
        </div>

    </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import {
    getManageList
} from '@/api/manage.js'
const chart = ref(null)
const echart = ref(null)
const rework = ref([])
const repair = ref([])
const parts = ref([])
const pass = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(100000)
const searchInfo = ref({})
function initChart() {
    getNcrCount();
}

const getNcrCount = async () => {
    const res = await getManageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (res.code === 0) {
        statistics_department(res.data.list)
        chart.value = echarts.init(echart.value)
        setOptions()
        document.addEventListener('resize', () => {
            chart.value?.resize()
        })
        return
    }
    ElMessage({
        type: 'error',
        message: 'Network error',
        showClose: true
    })
}

const statistics_department = (sd) => {
    sd.forEach(item => {
        switch (item.operation_type) {
            case '返工':
                rework.value.push(item.operation_type)
                break
            case '返修':
                repair.value.push(item.operation_type)
                break
            case '配做':
                parts.value.push(item.operation_type)
                break
            case '报废':
                pass.value.push(item.operation_type)
                break
            case '让步':
                pass.value.push(item.operation_type)
                break
            default:
                break
        }
    })
}


const setOptions = () => {
    chart.value.setOption({
        backgroundColor: 'transparent',
        grid: {
            left: '0',
            right: '0',
            top: '60%',
            bottom: '0',
            containLabel: false
        },
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'line',
                lineStyle: {
                    color: 'rgba(50, 216, 205, 1)'
                },
            }
        },
        xAxis: [{
            type: 'category',
            axisLine: {
                show: false
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                show: false
            }
        }],
        yAxis: [
            {
                splitLine: {
                    show: false,
                },
                axisLine: {
                    show: false
                },
                axisLabel: {
                    show: false,
                },
                axisTick: {
                    show: false
                }
            },

        ],
        series: [
            {
                name: '今日',
                type: 'line',
                smooth: true,
                stack: '总量',
                symbolSize: 5,
                showSymbol: false,
                itemStyle: {
                    normal: {
                        color: '#23D0C4',
                        lineStyle: {
                            color: '#23D0C4',
                            width: 1
                        },
                    }
                },
                areaStyle: {
                    normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                            {
                                offset: 0,
                                color: 'rgba(50, 216, 205, 1)'
                            },
                            {
                                offset: 1,
                                color: 'rgba(255, 255, 255, 0.2)'
                            }
                        ], false),
                    }
                },
                data: [rework.value.length, rework.value.length, repair.value.length, pass.value.length, parts.value.length]
            },
        ]
    })
}

onMounted(() => {
    nextTick(() => {
        initChart()
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
.ReclaimMileage-box {
    height: 120px;
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

    &-data {
        position: absolute;
        top: 0;
        left: 0;
        bottom: 10px;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 10;
        user-select: none;
        pointer-events: none;

        &-left {
            height: 60%;
            display: flex;
            align-items: flex-start;
            flex-direction: column;
            justify-content: center;

            div {
                color: #999;
                font-size: 12px;
                margin-bottom: 8px;
            }
        }

        &-center {
            height: 60%;
            border-right: 1px solid #eee;
            border-left: 1px solid #eee;
            display: flex;
            align-items: flex-start;
            flex-direction: column;
            justify-content: center;
            padding: 0 10px;

            div {
                color: #999;
                font-size: 12px;
                margin-bottom: 8px;
            }
        }

        &-right {
            height: 60%;
            padding-left: 10px;
            display: flex;
            align-items: flex-start;
            flex-direction: column;
            justify-content: center;

            div {
                color: #999;
                font-size: 12px;
                margin-bottom: 8px;
            }
        }
    }
}

.in-line {
    --color: #5BC2A4;
}

.out-line {
    --color: #DF534E;
}

.number {
    color: #1d1d1f !important;
    font-size: 18px !important;
    font-weight: 500;

    span {
        font-size: 12px;
        color: #DF534E;
        font-weight: 400;
        transform: scale(0.8333);
    }
}
</style>