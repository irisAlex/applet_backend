<template>
    <div style="text-align:center">
        <canvas ref="canvasRef" width="40" height="40" @click="handleClick"></canvas>
    </div>
</template>

<script setup>
import { ref, onMounted, nextTick, watch } from 'vue';
const canvasRef = ref(null);
var x = 20;
var y = 20;
var RADIUS = 10;
let circle = defineModel(1)
const props = defineProps(['defaultValue'])
circle.value = props.defaultValue
let number = 1
number = props.defaultValue
onMounted(() => {
    nextTick(() => {
        drawCircle();
    })
})

function drawCircle() {
    const canvas = canvasRef.value;
    const ctx = canvas.getContext('2d');

    // 清空画布
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.beginPath();
    switch (number) {
        case 1:
            ctx.arc(x, y, RADIUS, Math.PI / 2, Math.PI); //sAngle 90 ,eAngle 180    ctx.fillStyle = '#333';
            break;
        case 2:
            ctx.arc(x, y, RADIUS, Math.PI / 2, Math.PI * 1.5); //sAngle 90 ,eAngle 180    ctx.fillStyle = '#333';
            break
        case 3:
            ctx.arc(x, y, RADIUS, Math.PI / 2, Math.PI * 2); //sAngle 90 ,eAngle 180    ctx.fillStyle = '#333';
            break
        case 4:
            ctx.arc(x, y, RADIUS, Math.PI / 2, Math.PI * 3); //sAngle 90 ,eAngle 180    ctx.fillStyle = '#333';
            break
        default:
            break
    }
    ctx.lineTo(x, y);
    ctx.fill();
    ctx.restore();
    ctx.stroke();
}

function handleClick() {
    number++
    circle.value = number
    if (number > 4) {
        number = 1
        circle.value = number
    }
    drawCircle();

}

</script>

<style scoped></style>