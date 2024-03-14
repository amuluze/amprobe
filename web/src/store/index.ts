import { useUserStore } from '@/store/modules/user';

// 注册子模块
const useStore = () => ({
    user: useUserStore(),
});

export default useStore;
