<template>
  <div class="mt-2">
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <template #header>
            <el-divider>Alex</el-divider>
          </template>
        </el-card>
        <el-card style="margin-top: 20px">
          <template #header>
            <div>Alex团队</div>
          </template>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header>
            <div>提交记录</div>
          </template>
          <div>
            <el-timeline>
              <el-timeline-item v-for="(item, index) in dataTimeline" :key="index" :timestamp="item.from"
                placement="top">
                <el-card>
                  <h4>{{ item.title }}</h4>
                  <p>{{ item.message }}</p>
                </el-card>
              </el-timeline-item>
            </el-timeline>
          </div>
          <el-button class="load-more" type="primary" link @click="loadMore">
            Load more
          </el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Commits, Members } from '@/api/github'
import { formatTimeToStr } from '@/utils/date'
const page = ref(0)

defineOptions({
  name: 'About'
})

const loadMore = () => {
  page.value++
  loadCommits()
}

const dataTimeline = ref([])
const loadCommits = () => {
  Commits(page.value).then(({ data }) => {
    data.forEach((element) => {
      if (element.commit.message) {
        dataTimeline.value.push({
          from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        })
      }
    })
  })
}

const members = ref([])
const loadMembers = () => {
  Members().then(({ data }) => {
    members.value = data
    members.value.sort()
  })
}

loadCommits()
loadMembers()

</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}


.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
