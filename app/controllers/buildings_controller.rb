class BuildingsController < ApplicationController
  
  def index
    @items = Building.all.limit(50)
  end

  def show
    @building = Building.find_by_guid(params[:id])
  end

  def new
    @building = Building.new
  end

  def create
    @item = Building.new(item_params)

    if @item.save
      redirect_to building_path(@item.guid), notice: "was successfully created."
    else
      render :new, status: :unprocessable_entity
    end
  end

  private

  def item_params
    params.require(:building).permit(:address)
  end

end
