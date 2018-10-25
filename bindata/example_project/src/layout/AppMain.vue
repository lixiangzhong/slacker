<template>
    <section>
        <transition name="slide-fade" mode="out-in">
            <router-view></router-view>
        </transition>
    </section>
</template>

<script>
    export default {
      name: 'AppMain'
    }
</script>

<style scoped>
    .slide-fade-enter-active,
    .slide-fade-leave-active {
      transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
    }
    .slide-fade-enter,
    .slide-fade-leave-to {
      transform: translateX(10px);
      opacity: 0;
    }
</style>


