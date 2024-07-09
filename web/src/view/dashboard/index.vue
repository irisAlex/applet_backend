<template>
  <div>
    <div class="data-center-box">
      <CenterCard title="NCR数量">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ChainRatio />
        </template>
      </CenterCard>
      <CenterCard title="异常发现数">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ReclaimMileage />
        </template>
      </CenterCard>
      <CenterCard title="分类统计">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <Order />
        </template>
      </CenterCard>
      <CenterCard title="报废率">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <RecoveryRate />
        </template>
      </CenterCard>
      <CenterCard title="不合格品处理" style="grid-column-start: span 3;">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <lineCharts />
        </template>
      </CenterCard>
      <CenterCard title="统计" style="grid-auto-columns: 1.5fr">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <part />
        </template>
      </CenterCard>
    </div>
  </div>
</template>

<script setup>
import CenterCard from '@/view/dashboard/dashboardCard/centerCard.vue'
import ChainRatio from '@/view/dashboard/dashboardRatio/chainRatio.vue'
import ReclaimMileage from '@/view/dashboard/dashboardMileage/ReclaimMileage.vue'
import RecoveryRate from '@/view/dashboard/dashboardRate/RecoveryRate.vue'
import Order from '@/view/dashboard/dashboardOrder/order.vue'
import lineCharts from '@/view/dashboard/dashboardLineCharts/lineCharts.vue'
import part from '@/view/dashboard/dashboardPart/part.vue'

</script>

<script>
export default {
  name: 'DataCenter',
}
</script>

<style lang="scss" scoped>
.data-center-box {
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.5fr 2fr 1.5fr;
  column-gap: 10px;
}
</style>