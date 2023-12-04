/**
 * 看板数据
 * @cmd yao run scripts.dash.Data
 * @returns
 */
function Data() {
  var obj = Process("flows.stat.data");
  console.log(obj);
  return obj;
}
