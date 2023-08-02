declare global {
    interface String {
      // 在这里定义你要增加的方法
      replaceMultipleNbsp(): string;
      replaceMultipleBR(): string;
      removeTagP(): string;
      splitStringByBr(): string[];
    }
  }

// 将string中的 nbsp 换成空格
String.prototype.replaceMultipleNbsp = function () {
    const regex = /&nbsp;/g;
    return this.replace(regex, ' ');
};

String.prototype.replaceMultipleBR = function () {
    const regex = /(<br\s*\/?>)\s*(<br\s*\/?>)+/gi;
    return this.replace(regex, '<br>');
};

/**
 * 去除字符串中的 <p> 和 </p> 标签
 */
String.prototype.removeTagP = function () {
    const regex = /<p>/g;
    return this.replace(regex, '<br>');
};

/**
 * 将字符串按照 <br> 分割为字符串数组 空字符串不计入数组中
 */
String.prototype.splitStringByBr = function () {
    return this.split('\n').filter(str => str.trim() !== '');
};
// 将传入的字符串 str 中的四个空格符前插入换行符
export function insertNewlineBeforeFourSpaces(str: string): string {
    return str.replaceMultipleNbsp().replaceMultipleBR().replace(/ {4}/g, '\n');
}
export function splitStringByBr(str: string | undefined): string[] {
    if (!str) return [];
    return str.splitStringByBr();
}