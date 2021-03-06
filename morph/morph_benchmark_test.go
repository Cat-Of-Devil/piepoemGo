// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package morph

import "testing"

var benchWords = []string{
	"абдулхаковичам",
	"авентине",
	"анзиях",
	"анолиевичу",
	"ассенизирующий",
	"афтению",
	"ашимовичах",
	"аэрогидромеханический",
	"балыкиной",
	"бающегося",
	"белизцами",
	"бомбежная",
	"бравшись",
	"бурундийкам",
	"бычинам",
	"вбивающихся",
	"венегдитович",
	"взбудораживаемся",
	"вздуванью",
	"взрывостоек",
	"видеодокументах",
	"викторасов",
	"владасовна",
	"вогнуть",
	"водорослевидном",
	"воронова",
	"вспышкою",
	"выгонною",
	"выгрязнилась",
	"вылуживаема",
	"высасываете",
	"высуженного",
	"выясняющие",
	"гедонистическим",
	"гелиевой",
	"геращенков",
	"гильотинирующему",
	"голосистое",
	"голышмановское",
	"гомопаузе",
	"гопчиков",
	"горочна",
	"горянками",
	"грабленому",
	"грохольское",
	"гуртов",
	"даутов",
	"двадцатикомпонентными",
	"двузначность",
	"депрессанте",
	"деснах",
	"джабраиловичи",
	"диатермокоагуляцию",
	"додержавшим",
	"доопределите",
	"дорубавшие",
	"доукомплектовывавшиеся",
	"дробные",
	"ерванды",
	"живописцем",
	"журфикса",
	"заборонило",
	"завозчиками",
	"задобрите",
	"зазимкам",
	"заливное",
	"замедлившегося",
	"занашивающимся",
	"запячивавшейся",
	"затасканном",
	"затаскивавшую",
	"заточалось",
	"извозилась",
	"изолирующимися",
	"инфицировавшемся",
	"исклевываю",
	"карагалинки",
	"караибам",
	"клиопатр",
	"колачивавшегося",
	"колдуя",
	"комбинаторику",
	"компаративизму",
	"корнфлексе",
	"крапивинская",
	"краскотеры",
	"кроншнепом",
	"крупитчатой",
	"лесбиянок",
	"липни",
	"лихорадимою",
	"лицемерном",
	"локусе",
	"маврами",
	"молнирующей",
	"монегасков",
	"московитах",
	"мяукая",
	"наваливаемых",
	"навинтившийся",
	"навинчивая",
	"надклепавшею",
	"надковывайся",
	"накарябали",
	"накоксованной",
	"напугавшейся",
	"натискивавшегося",
	"натравляемых",
	"небезупречном",
	"неделегатский",
	"нежелчны",
	"незагустелую",
	"незрелыми",
	"неизвинительной",
	"неизодранное",
	"неоспоримостям",
	"непиратское",
	"непревратимые",
	"несредиземноморских",
	"нефтемаслозаводе",
	"новоярославскими",
	"обветревшие",
	"обесчеловечиваемые",
	"обкусавшему",
	"облюбившему",
	"обстрачивавший",
	"обшмонайте",
	"овиваешь",
	"окармливавшиеся",
	"омертвлявшемся",
	"оповещающейся",
	"опробовавшейся",
	"ортодоксами",
	"осведомлявшего",
	"остеитами",
	"отабековичам",
	"отакелаживающемуся",
	"отгрызающимися",
	"отданного",
	"отдельщицах",
	"откладывайте",
	"отполировывающее",
	"отучившаяся",
	"отшпандоривавшуюся",
	"охаживавшейся",
	"ошибавшимися",
	"перверсиями",
	"перевинтятся",
	"переделяем",
	"передислоцировало",
	"пережиравшиеся",
	"перемогавшей",
	"перемоточные",
	"перехихикнувшихся",
	"поворованном",
	"подвыросло",
	"подглядывайте",
	"подключит",
	"подлеченными",
	"подтягивавшейся",
	"подцентр",
	"поземельная",
	"покидываемыми",
	"попроходимее",
	"попрятал",
	"поразогнавшим",
	"пороховниц",
	"послевузовское",
	"похороненный",
	"поченикина",
	"почердачнее",
	"почнемте",
	"предустановившем",
	"приберегать",
	"прикрикни",
	"примежевываемо",
	"приморскагропромэнерго",
	"припоминавшиеся",
	"притуливший",
	"притупленная",
	"приумявшего",
	"проваживает",
	"прогрузили",
	"променады",
	"пропиралось",
	"пропитье",
	"проплясанною",
	"простонародна",
	"простригающей",
	"протаскавшееся",
	"протранжирено",
	"профильтровался",
	"прошлифованы",
	"психофармакологическими",
	"пукальщика",
	"радиобиотелеметрия",
	"разархивировал",
	"разбраковывайте",
	"развалившая",
	"разворачивавшему",
	"размечаемые",
	"разнуздалось",
	"разрисовавшего",
	"раковиной",
	"расселений",
	"расщемив",
	"рахиловнами",
	"рекордизмы",
	"реорганизовываемы",
	"ризвановной",
	"свариваемого",
	"светленным",
	"сдвурушничало",
	"севмаша",
	"силишку",
	"силуэтными",
	"сказаний",
	"сказуемость",
	"скоробилась",
	"слогоделению",
	"снисходительным",
	"спилбергом",
	"сплавившею",
	"среднерусской",
	"сроднивши",
	"стайерская",
	"сушимое",
	"сцарапывавшим",
	"терпеливое",
	"титаноносных",
	"тифлит",
	"томатном",
	"топливоэнергоснабжающее",
	"трамонтана",
	"трубокладе",
	"увлажило",
	"увлекся",
	"углубляющей",
	"усложняйтесь",
	"услужит",
	"утеплительною",
	"ушераздирающа",
	"фитоценологиях",
	"фишки",
	"хапки",
	"хоронивший",
	"цедентом",
	"целить",
	"цилиндрового",
	"широкобедрую",
	"щаженных",
	"щекотавшую",
	"экранизированы",
	"янисовна",
	"яношевнах",
	"ярополковной",
}

var tmp []string

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp, _, _ = Parse(benchWords[i%len(benchWords)])
	}
}
