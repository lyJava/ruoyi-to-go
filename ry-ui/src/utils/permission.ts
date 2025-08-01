import useUserStore from '@/store/modules/user'

/**
 * 字符权限校验
 * 
 * @param {Array} value 校验值
 * @returns {Boolean}
 */
export const checkPermi = (value: string | any[]) => {
    if (value && value instanceof Array && value.length > 0) {
        const permissions = useUserStore().permissions;
        const permissionDatas = value;
        const all_permission = "*:*:*";

        const hasPermission = permissions.some((permission: string) => {
            return (
                all_permission === permission || permissionDatas.includes(permission)
            );
        });

        if (!hasPermission) {
            return false;
        }
        return true;
    } else {
        console.error(
            `need roles! Like checkPermi="['system:user:add','system:user:edit']"`
        );
        return false;
    }
};

/**
 * 角色权限校验
 * 
 * @param {Array} value 校验值
 * @returns {Boolean}
 */
export const checkRole = (value: string | any[]): boolean => {
    if (value && value instanceof Array && value.length > 0) {
        const roles = useUserStore().roles;
        const permissionRoles = value;
        const super_admin = "admin";

        const hasRole = roles.some((role: string) => {
            return super_admin === role || permissionRoles.includes(role);
        });

        if (!hasRole) {
            return false;
        }
        return true;
    } else {
        console.error(`need roles! Like checkRole="['admin','editor']"`);
        return false;
    }
};
